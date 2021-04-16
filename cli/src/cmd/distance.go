package cmd

import (
	"github.com/lolPants/redts/cli/src/pkg/edts"
	"github.com/spf13/cobra"
)

var (
	distanceCmd = &cobra.Command{
		Use:                "distance",
		Short:              "finds the distance between two or more systems",
		Args:               cobra.ArbitraryArgs,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			edts.RunCommand(cmd.Use, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(distanceCmd)
}
