package domain

import (
	"github.com/BurntSushi/toml"
	"github.com/nlopes/slack"
)

// Env should have config and rtm
type Env struct {
	Config Config
	rtm    *slack.RTM
}

// Config should have Slack Config
type Config struct {
	Token string
	Botid string
}

// NewEnviroment return Config and rtm
func NewEnviroment(configfile string) *Env {

	var config Config

	_, err := toml.DecodeFile(configfile, &config)
	if err != nil {
		// Error Handling
	}

	slackrtm := slack.New(config.Token).NewRTM()

	return &Env{
		Config: config,
		rtm:    slackrtm,
	}
}
