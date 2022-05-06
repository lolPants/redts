package config

import (
	"fmt"
	"os"
)

func Load() *Config {
	cfg := Config{}
	err := cfg.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read config")
		os.Exit(1)
	}

	return &cfg
}
