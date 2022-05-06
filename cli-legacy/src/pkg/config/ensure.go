package config

import (
	"fmt"
	"os"
)

func (c *Config) EnsureRequiredFields() bool {
	if c.URL == "" {
		fmt.Fprintln(os.Stderr, "error: no remote url set")
		fmt.Fprintln(os.Stderr, "use: redts config set url <url>")

		return false
	}

	return true
}

func (c *Config) EnsureAuthFields() bool {
	if c.Username == "" {
		fmt.Fprintln(os.Stderr, "error: no username")
		fmt.Fprintln(os.Stderr, "use: redts config set username <username>")

		return false
	}

	if c.Token == "" {
		fmt.Fprintln(os.Stderr, "error: no auth token set")
		fmt.Fprintln(os.Stderr, "use: redts config set token <token>")

		return false
	}

	return true
}
