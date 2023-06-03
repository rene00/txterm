package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func accountMatchCmd(cli *cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "account-match",
	}

	return cmd
}

func createAccountMatchCmd(cli *cli) *cobra.Command {
	var flags struct {
		Name        string
		Description string
		AccountID   int
	}
	fmt.Println(flags)
	return nil
}
