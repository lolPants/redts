package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/lolPants/redts/cli/src/pkg/config"
	"github.com/spf13/cobra"
)

var (
	configSetCmd = &cobra.Command{
		Use:   "set",
		Short: "set config key value",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			key := strings.ToLower(args[0])
			value := args[1]
			cfg := config.Load()

			validKey := cfg.HasField(key)
			if !validKey {
				fmt.Fprintf(os.Stderr, "error: invalid config key `%s`\n", key)
				os.Exit(1)

				return
			}

			if key == "url" {
				value = strings.TrimSuffix(value, "/")
			}

			cfg.SetField(key, value)
			err := cfg.Save()
			if err != nil {
				panic(err)
			}
		},
	}
)

func init() {
	configCmd.AddCommand(configSetCmd)
}
