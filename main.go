package main

import (
	"flag"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zvlb/release-watcher/internal/config"
	"github.com/zvlb/release-watcher/internal/providers"
	"github.com/zvlb/release-watcher/internal/providers/github"
	"github.com/zvlb/release-watcher/internal/recievers"
	"github.com/zvlb/release-watcher/internal/recievers/telegram"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config-file", "/release-watcher/config.yaml", "Config file.")
	flag.Parse()

	cfg, err := config.GetConfiguration(configFile)
	if err != nil {
		panic(err)
	}

	err = StartWatchers(cfg)
	if err != nil {
		panic(err)
	}
}

func StartWatchers(cfg config.Config) error {
	recievers := generateRecievers(cfg)
	providers, err := generateProviders(cfg)
	if err != nil {
		return err
	}

	// Start for GitHub Providers
	for _, p := range providers {
		go Watcher(p, recievers)
	}

	for {
		time.Sleep(time.Hour * 100)
	}
}

func Watcher(provired providers.Provider, recievers []recievers.Reciever) error {
	log.Infof("Start Release Watch for %v", provired.GetName())
	for {
		title, description, link, err := provired.WatchReleases()
		if err != nil {
			return err
		}

		log.Infof("Find New Release for %v", provired.GetName())
		for _, r := range recievers {
			r.SendData(title, description, link)
		}
	}
}

func generateRecievers(cfg config.Config) []recievers.Reciever {
	recievers := []recievers.Reciever{}

	// Get telegram Recievers
	for _, t := range cfg.Recievers.Telegram {
		r := telegram.New(t.Token, t.ChatID)
		recievers = append(recievers, r)
	}

	return recievers
}

func generateProviders(cfg config.Config) ([]providers.Provider, error) {
	providers := []providers.Provider{}

	// Get github Provides
	for _, g := range cfg.Releases.Github {
		p, err := github.New(g, &http.Client{})
		if err != nil {
			return providers, err
		}

		providers = append(providers, p)
	}

	return providers, nil
}
