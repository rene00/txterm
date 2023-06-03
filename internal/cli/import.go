package cli

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"txterm/internal/store"

	"github.com/aclindsa/ofxgo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rene00/gnucash-sqlboiler/gnucash"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			str, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}

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
			for _, gnucashAccount := range gnucashAccounts {
				account, err := str.CreateAccount(cmd.Context(), gnucashAccount.Name, gnucashAccount.Description.String, strings.ToLower(gnucashAccount.AccountType), 0)
				if err != nil {
					return fmt.Errorf("failed to create account: %w", err)
				}
				fmt.Println(account)
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
			str, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}

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

				imprt, err := str.SaveImport(cmd.Context(), filepath.Base(f.Name()), stmt.BalAmt.Rat, stmt.DtAsOf.Time)
				if err != nil {
					return err
				}

				for _, tran := range stmt.BankTranList.Transactions {
					_, err := str.SaveTransaction(cmd.Context(), tran.DtPosted.Time, tran.Memo.String(), tran.TrnAmt.Rat, *imprt)
					var transactionError *store.TransactionError
					switch {
					case errors.As(err, &transactionError):
						if transactionError.Code() == store.TransactionErrDuplicate {
							continue
						}
						return err
					case err != nil:
						return err
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
