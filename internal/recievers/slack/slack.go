package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/zvlb/release-watcher/internal/config"
	"github.com/zvlb/release-watcher/internal/recievers"
)

var (
	slackAPI      = ""
	slackNo200Err = errors.New("status code for request to Slack API is not 200")
)

type SlackReciever struct {
	ChannelName string `yaml:"channelName"`
	Hook        string `yaml:"hook"`
	Config      config.Config
}

func New(channelname, hook string, config config.Config) recievers.Reciever {
	return &SlackReciever{
		ChannelName: channelname,
		Hook:        hook,
		Config:      config,
	}
}

func (sr *SlackReciever) GetName() string {
	return fmt.Sprintf("Slack channel %s ", sr.ChannelName)
}

func (sr *SlackReciever) SendData(name, release, description, link string) error {
	url := fmt.Sprintf("%v", sr.Hook) //??????

	var text string
	if sr.Config.Params.SendReleaseDescription {
		text = fmt.Sprintf("%v. Release: %v\n%v\n\n%v", name, release, description, link)
	} else {
		text = fmt.Sprintf("%v. Release: %v\n%v", name, release, link)
	}

	body, err := json.Marshal(map[string]string{
		"text":    text,
		"channel": sr.ChannelName,
	})

	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if response.StatusCode != http.StatusOK {
		return slackNo200Err
	}

	return nil
}
