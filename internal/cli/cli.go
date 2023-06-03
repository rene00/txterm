package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type cli struct {
	debug bool
}

func Execute() {
	cli := &cli{}
	rootCmd := &cobra.Command{
		Use: "txterm",
	}
	rootCmd.PersistentFlags().BoolVar(&cli.debug, "debug", false, "Enable debug")

	rootCmd.AddCommand(importCmd(cli))

	if err := rootCmd.ExecuteContext(context.TODO()); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
