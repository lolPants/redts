package cmd

import (
	"os"
	"strings"

	"github.com/lolPants/redts/cli/src/pkg/config"
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
			cfg := config.Load()
			if !cfg.EnsureRequiredFields() {
				os.Exit(1)
			}

			var parsedArgs strings.Builder
			for i, arg := range args {
				parsedArgs.WriteRune('"')
				parsedArgs.WriteString(arg)
				parsedArgs.WriteRune('"')

				if i+1 != len(args) {
					parsedArgs.WriteRune(' ')
				}
			}

			err := edts.CallAPI(cfg, "edts", parsedArgs.String())
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(edtsCmd)
}
