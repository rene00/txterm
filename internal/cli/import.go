package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"txterm/internal/store"

	"github.com/aclindsa/ofxgo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func importCmd(cli *cli) *cobra.Command {
	var flags struct {
		File string
	}

	var cmd = &cobra.Command{
		Use: "import",
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
