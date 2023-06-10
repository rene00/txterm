package cli

import (
	"fmt"
	"net/http"
	"txterm/internal/store"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			str, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}

			transactions, err := str.ListTransactions(cmd.Context())
			if err != nil {
				return fmt.Errorf("list transactions: %w", err)
			}

			accounts, err := str.ListAccounts(cmd.Context())
			if err != nil {
				return fmt.Errorf("list accounts: %w", err)
			}

			r := gin.Default()
			r.LoadHTMLGlob("templates/*.html")

			r.GET("/", func(c *gin.Context) {
				c.HTML(http.StatusOK, "index.html", gin.H{"transactions": transactions})
			})

			r.GET("/accounts", func(c *gin.Context) {
				c.HTML(http.StatusOK, "accounts.html", gin.H{"accounts": accounts})
			})

			r.Run(fmt.Sprintf(":%d", flags.Port))
			return nil
		},
	}

	cmd.Flags().IntVar(&flags.Port, "port", 8000, "Port to listen on")
	return cmd
}
