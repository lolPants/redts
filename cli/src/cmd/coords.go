package cmd

import (
	"github.com/lolPants/redts/cli/src/pkg/edts"
	"github.com/spf13/cobra"
)

var (
	coordsCmd = &cobra.Command{
		Use:                "coords",
		Short:              "returns the coordinates of given systems",
		Args:               cobra.ArbitraryArgs,
		DisableFlagParsing: true,
		Run: func(cmd *cobra.Command, args []string) {
			edts.RunCommand(cmd.Use, args)
		},
	}
)

func init() {
	rootCmd.AddCommand(coordsCmd)
}
