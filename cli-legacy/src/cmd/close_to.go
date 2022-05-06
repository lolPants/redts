package cmd

import (
	"github.com/lolPants/redts/cli/src/pkg/edts"
	"github.com/spf13/cobra"
)

var (
	closetoCmd = &cobra.Command{
		Use:                "close_to",
		Short:              "finds systems close to others, optionally with constraints",
		Args:               cobra.ArbitraryArgs,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			edts.RunCommand(cmd.Use, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(closetoCmd)
}
