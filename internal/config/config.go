package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Releases  Releases  `yaml:"releases"`
	Recievers Recievers `yaml:"recievers"`
	Params    Params    `yaml:params`
}

type Releases struct {
	Github []string `yaml:"github"`
}

type Recievers struct {
	Telegram []TelegramReciever `yaml:"telegram"`
	Slack    []SlackReciever    `yaml:"slack"`
}

type SlackReciever struct {
	ChannelName string `yaml:"channelName"`
	Hook        string `yaml:"hook"`
}

type TelegramReciever struct {
	ChatID string `yaml:"chatID"`
	Token  string `yaml:"token"`
}

type Params struct {
	SendReleaseDescription bool `yaml:"sendReleaseDescription"`
}

func GetConfiguration(path string) (Config, error) {
	var config Config

	content, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
