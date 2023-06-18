package cli

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"txterm/db/model"
	"txterm/internal/logwrap"
	"txterm/internal/store"

	"github.com/aclindsa/ofxgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rene00/gnucash-sqlboiler/gnucash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func importCmd(cli *cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "import",
	}

	cmd.AddCommand(importTransactionsCmd(cli))
	cmd.AddCommand(importAccountsCmd(cli))
	return cmd
}

func importAccountsCmd(cli *cli) *cobra.Command {
	var flags struct {
		GnuCashURI string
	}

	var cmd = &cobra.Command{
		Use: "accounts",
		PreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindPFlag("gnucash-uri", cmd.Flags().Lookup("gnucash-uri"))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logwrap.New("logger", os.Stdout, false)
			logger.SetLevel(logwrap.INFO)
			if cli.debug {
				logger.SetLevel(logwrap.DEBUG)
			}

			s, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}
			defer s.DB.Close()
			boil.SetDB(s.DB)

			gnucashDB, err := sql.Open("sqlite3", flags.GnuCashURI)
			if err != nil {
				return fmt.Errorf("failed to open gnucash sqlite3: %w", err)
			}
			defer gnucashDB.Close()
			boil.SetDB(gnucashDB)

			gnucashAccounts, err := gnucash.Accounts(qm.Where("account_type!=?", "ROOT")).All(cmd.Context(), gnucashDB)
			if err != nil {
				return fmt.Errorf("failed to find gnucash accounts: %w", err)

			}
			for _, gcAccount := range gnucashAccounts {
				accountType, err := model.AccountTypes(qm.Where("name=?", strings.ToLower(gcAccount.AccountType))).One(cmd.Context(), s.DB)
				switch {
				case errors.Is(err, sql.ErrNoRows):
					accountType = &model.AccountType{Name: strings.ToLower(gcAccount.AccountType)}
					if err := accountType.Insert(cmd.Context(), s.DB, boil.Infer()); err != nil {
						return fmt.Errorf("insert accounttype: %w", err)
					}
				case err != nil:
					return fmt.Errorf("account types: %w", err)
				}

				account := model.Account{
					Name:          gcAccount.Name,
					Description:   null.StringFrom(gcAccount.Description.String),
					AccountTypeID: accountType.ID,
					Code:          null.StringFrom(gcAccount.Code.String),
				}

				if err := account.Insert(cmd.Context(), s.DB, boil.Infer()); err != nil {
					return fmt.Errorf("account: %w", err)
				}
				logger.Debug(fmt.Sprintf("created account: %s", account.Name))
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.GnuCashURI, "gnucash-uri", "", "Gnucash URI")
	return cmd
}

func importTransactionsCmd(cli *cli) *cobra.Command {
	var flags struct {
		File string
	}

	var cmd = &cobra.Command{
		Use: "transactions",
		PreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindPFlag("file", cmd.Flags().Lookup("file"))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logwrap.New("logger", os.Stdout, false)
			logger.SetLevel(logwrap.INFO)
			if cli.debug {
				logger.SetLevel(logwrap.DEBUG)
			}

			s, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}
			defer s.DB.Close()
			boil.SetDB(s.DB)

			f, err := os.Open(flags.File)
			if err != nil {
				return err
			}
			defer f.Close()

			response, err := ofxgo.ParseResponse(f)
			if err != nil {
				return err
			}

			if stmt, ok := response.CreditCard[0].(*ofxgo.CCStatementResponse); ok {
				// Find account with code which has accountID. This will be
				// used as the source account.
				sourceAccount, err := model.Accounts(
					[]qm.QueryMod{
						qm.Where("code=?", stmt.CCAcctFrom.AcctID.String()),
						qm.InnerJoin("account_type at on account.account_type_id = at.id"),
						qm.Where("at.name=?", "liability"),
					}...,
				).One(cmd.Context(), s.DB)
				if err != nil {
					return fmt.Errorf("source account: %w", err)
				}

				importRun := model.ImportRun{
					DateCreated: time.Now(),
					Filename:    filepath.Base(f.Name()),
				}
				if err := importRun.Insert(cmd.Context(), s.DB, boil.Infer()); err != nil {
					return fmt.Errorf("insert import run: %w", err)
				}

				for _, t := range stmt.BankTranList.Transactions {
					tx, err := s.DB.BeginTx(cmd.Context(), nil)
					if err != nil {
						return fmt.Errorf("begin tx: %w", err)
					}
					defer tx.Rollback()

					transaction := model.Transaction{
						DateCreated: t.DtPosted.Time,
						DatePosted:  time.Now(),
						Memo:        t.Memo.String(),
					}
					if err := transaction.Insert(cmd.Context(), tx, boil.Infer()); err != nil {
						return fmt.Errorf("insert tx: %w", err)
					}

					transactionImportRun := model.TransactionImportRun{
						TransactionID: transaction.ID,
						ImportRunID:   importRun.ID,
					}
					if err := transactionImportRun.Insert(cmd.Context(), tx, boil.Infer()); err != nil {
						return fmt.Errorf("insert transaction accountrun: %w", err)
					}

					for _, split := range []model.Split{
						{
							TransactionID: transaction.ID,
							AccountID:     sourceAccount.ID,
							ValueNum:      t.TrnAmt.Num().Int64(),
							ValueDenom:    t.TrnAmt.Denom().Int64(),
						},
						{
							TransactionID: transaction.ID,
						},
					} {
						if err := split.Insert(cmd.Context(), tx, boil.Infer()); err != nil {
							return fmt.Errorf("insert split: %w", err)
						}
					}

					if err := tx.Commit(); err != nil {
						return fmt.Errorf("commit: %w", err)
					}
				}
			} else {
				return fmt.Errorf("file not supported")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.File, "file", "", "File to import")
	return cmd
}
