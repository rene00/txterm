package cli

import (
	"fmt"
	"net/http"
	"txterm/db/model"
	"txterm/internal/modext"
	"txterm/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func serveCmd(cli *cli) *cobra.Command {
	var flags struct {
		Port int
	}

	var cmd = &cobra.Command{
		Use: "serve",
		PreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindPFlag("port", cmd.Flags().Lookup("port"))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}

			transactions, err := model.Transactions(
				[]qm.QueryMod{
					qm.Load(qm.Rels(model.TransactionRels.Splits, model.SplitRels.Account, model.AccountRels.AccountType)),
				}...).All(cmd.Context(), s.DB)
			if err != nil {
				return fmt.Errorf("all transactions: %w", err)
			}

			modextTransactions := []modext.Transaction{}
			for _, transaction := range transactions {
				modextTransactions = append(modextTransactions, modext.NewTransaction(*transaction))
			}

			accounts, err := model.Accounts().All(cmd.Context(), s.DB)
			if err != nil {
				return fmt.Errorf("all accounts: %w", err)
			}

			r := gin.Default()
			r.LoadHTMLGlob("templates/*.html")

			r.GET("/", func(c *gin.Context) {
				c.HTML(http.StatusOK, "index.html", gin.H{"transactions": modextTransactions, "accounts": accounts})
			})

			r.GET("/accounts", func(c *gin.Context) {
				c.HTML(http.StatusOK, "accounts.html", gin.H{"accounts": accounts})
			})

			r.Run(fmt.Sprintf("127.0.0.1:%d", flags.Port))
			return nil
		},
	}

	cmd.Flags().IntVar(&flags.Port, "port", 8000, "Port to listen on")
	return cmd
}
