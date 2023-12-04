package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Releases  Releases  `yaml:"releases"`
	Recievers Recievers `yaml:"recievers"`
}

type Releases struct {
	Github []string `yaml:"github"`
}

type Recievers struct {
	Telegram []TelegramReciever `yaml:"telegram"`
}

type TelegramReciever struct {
	ChatID string `yaml:"chatID"`
	Token  string `yaml:"token"`
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
