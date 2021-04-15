package cmd

import (
	"github.com/spf13/cobra"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "get/set config keys",
		Args:  cobra.ArbitraryArgs,
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
}
