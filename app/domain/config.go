package domain

import (
	"log"

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
		log.Printf("[Error] cannot find config.toml")
	}

	log.Printf(config.Token)

	return &config
}
