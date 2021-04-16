package config

type Config struct {
	URL      string `config:"url" toml:"url"`
	Username string `config:"username" toml:"username"`
	Token    string `config:"token" toml:"token"`
}
