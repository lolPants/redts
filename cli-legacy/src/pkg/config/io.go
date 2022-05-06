package config

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/pelletier/go-toml"
)

func (c *Config) filePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(home, ".config", "redts.toml"), nil
}

func (c *Config) Save() error {
	bytes, err := toml.Marshal(*c)
	if err != nil {
		return err
	}

	filepath, err := c.filePath()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, bytes, 0o644)
}

func (c *Config) Load() error {
	filepath, err := c.filePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		err = c.Save()
		if err != nil {
			return err
		}

		return nil
	}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	return toml.Unmarshal(bytes, c)
}
