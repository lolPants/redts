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

			modified := false
			cfg := config.Load()

			if key == "url" {
				suffix := "/"
				if strings.HasSuffix(value, suffix) {
					value = value[:len(value)-len(suffix)]
				}

				cfg.URL = value
				modified = true
			}

			if key == "username" {
				cfg.Username = value
				modified = true
			}

			if key == "token" {
				cfg.Token = value
				modified = true
			}

			if modified {
				err := cfg.Save()
				if err != nil {
					panic(err)
				}

				return
			}

			fmt.Fprintf(os.Stderr, "error: invalid config key `%s`\n", key)
			os.Exit(1)
		},
	}
)

func init() {
	configCmd.AddCommand(configSetCmd)
}
