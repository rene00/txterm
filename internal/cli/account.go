package cli

import (
	"fmt"
	"txterm/internal/store"

	"github.com/spf13/cobra"
)

func accountCmd(cli *cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "account",
	}
	return cmd
}

func listAccountCmd(cli *cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			str, err := store.New()
			if err != nil {
				return fmt.Errorf("new store: %w", err)
			}
			fmt.Println(str)
			return nil
		},
	}
	return cmd
}
