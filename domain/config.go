package domain

import (
	"github.com/BurntSushi/toml"
)

// Config should have Slack Config
type Config struct {
	Token string
	Botid string
}

// NewEnviroment return Config and rtm
func NewEnviroment() *Config {

	var config Config

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		// Error Handling
	}

	return &config
}
