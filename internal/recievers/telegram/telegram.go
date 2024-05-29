package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zvlb/release-watcher/internal/config"
	"github.com/zvlb/release-watcher/internal/recievers"
)

var (
	telegramAPI = "https://api.telegram.org/bot"
)

type TelegramReciever struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chatID"`
	Config config.Config
}

func New(token, chatID string, config config.Config) recievers.Reciever {
	return &TelegramReciever{
		Token:  token,
		ChatID: chatID,
		Config: config,
	}
}

func (tr *TelegramReciever) GetName() string {
	return fmt.Sprintf("Telegram chat %s ", tr.ChatID)
}

func (tr *TelegramReciever) SendData(name, release, description, link string) error {
	url := fmt.Sprintf("%v%v/%v", telegramAPI, tr.Token, "sendMessage")

	var text string
	if tr.Config.Params.SendReleaseDescription {
		text = fmt.Sprintf("<b>%v</b>. Release: <b>%v</b>\n%v\n\n%v", name, release, description, link)
	} else {
		text = fmt.Sprintf("<b>%v</b>. Release: <b>%v</b>\n%v", name, release, link)
	}

	body, err := json.Marshal(map[string]string{
		"chat_id":    tr.ChatID,
		"text":       text,
		"parse_mode": "HTML",
	})
	if err != nil {
		return err
	}

	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return no200Err
	}

	return nil
}
