package cmd

import (
	"github.com/lolPants/redts/cli/src/pkg/edts"
	"github.com/spf13/cobra"
)

var (
	findCmd = &cobra.Command{
		Use:                "find",
		Short:              "searches for systems and stations by name, including wildcards",
		Args:               cobra.ArbitraryArgs,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			edts.RunCommand(cmd.Use, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(findCmd)
}
