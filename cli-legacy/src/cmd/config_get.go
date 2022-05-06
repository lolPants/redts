package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/lolPants/redts/cli/src/pkg/config"
	"github.com/spf13/cobra"
)

var (
	configGetCmd = &cobra.Command{
		Use:   "get",
		Short: "get config key",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			key := strings.ToLower(args[0])
			cfg := config.Load()

			validKey := cfg.HasField(key)
			if !validKey {
				fmt.Fprintf(os.Stderr, "error: invalid config key `%s`\n", key)
				os.Exit(1)

				return
			}

			value := cfg.GetField(key)
			fmt.Println(value)
		},
	}
)

func init() {
	configCmd.AddCommand(configGetCmd)
}
