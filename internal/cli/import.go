package cli

import (
	"fmt"
	"os"
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
			fmt.Println(str)

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
				fmt.Printf("Balance: %s %s (as of %s)\n", stmt.BalAmt, stmt.CurDef, stmt.DtAsOf)
				for _, tran := range stmt.BankTranList.Transactions {
					fmt.Printf("%s %-15s [TrnType:%s] [Name:%s] [ExtdName:%s] [Memo:%s]\n", tran.DtPosted.Time.Format("2006-01-02"), tran.TrnAmt.String(), tran.TrnType, tran.Name, tran.ExtdName, tran.Memo)
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&flags.File, "file", "", "File to import")
	return cmd
}
