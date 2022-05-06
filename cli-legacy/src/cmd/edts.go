package cmd

import (
	"github.com/lolPants/redts/cli/src/pkg/edts"
	"github.com/spf13/cobra"
)

var (
	edtsCmd = &cobra.Command{
		Use:                "edts",
		Short:              "finds the optimal order to visit a set of stations, and can produce full routes between systems",
		Args:               cobra.ArbitraryArgs,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			edts.RunCommand(cmd.Use, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(edtsCmd)
}
