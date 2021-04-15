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
			cfg := config.Config{}
			err := cfg.Load()
			if err != nil {
				fmt.Fprintln(os.Stderr, "failed to read config")
				os.Exit(1)
			}

			if key == "url" {
				if cfg.URL == "" {
					fmt.Fprintln(os.Stderr, "error: url is unset")
					os.Exit(1)
				}

				fmt.Printf("%s\n", cfg.URL)
				return
			}

			if key == "username" {
				if cfg.URL == "" {
					fmt.Fprintln(os.Stderr, "error: username is unset")
					os.Exit(1)
				}

				fmt.Printf("%s\n", cfg.Username)
				return
			}

			if key == "token" {
				if cfg.URL == "" {
					fmt.Fprintln(os.Stderr, "error: token is unset")
					os.Exit(1)
				}

				fmt.Printf("%s\n", cfg.Token)
				return
			}

			fmt.Fprintf(os.Stderr, "error: invalid config key `%s`\n", key)
			os.Exit(1)
		},
	}
)

func init() {
	configCmd.AddCommand(configGetCmd)
}
