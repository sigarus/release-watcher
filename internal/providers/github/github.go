package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zvlb/release-watcher/internal/providers"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{}
}

var (
	// Github Api
	githubAPI_SCHEME = "https"
	githubAPI_URL    = "api.github.com/repos"
	// URI for get last release for repository
	lastReleaseReq = "releases/latest"

	// CLient
	Client HTTPClient
)

type GithubProvider struct {
	Path    string `yaml:"path"`
	Release ReleaseInfo
	client  *http.Client
}

func New(path string, client *http.Client) (providers.Provider, error) {
	gp := GithubProvider{
		Path:   path,
		client: client,
	}

	if gp.client == nil {
		gp.client = &http.Client{}
	}

	// Get actual release
	release, err := gp.getRelease()
	if err != nil {
		return nil, err
	}

	gp.Release = release

	return &gp, nil
}

func (gp *GithubProvider) WatchReleases() (title, description, link string, err error) {
	gp.Release, err = gp.getRelease()
	if err != nil {
		return
	}

	for {
		newReleaseExist := false
		newReleaseExist, err = gp.newReleaseExist()
		if err != nil {
			return
		}

		if newReleaseExist {
			return gp.getTitle(), gp.Release.Body, gp.Release.HtmlUrl, nil
		}

		time.Sleep(10 * time.Minute)
	}
}

func (gp *GithubProvider) GetName() string {
	return gp.Path
}

func (gp *GithubProvider) newReleaseExist() (bool, error) {
	newRelease, err := gp.getRelease()
	if err != nil {
		return false, err
	}

	if newRelease.TagName != gp.Release.TagName {
		gp.updateRelease(newRelease)
		return true, nil
	}
	return false, nil
}

func (gp *GithubProvider) getRelease() (ReleaseInfo, error) {
	var ri ReleaseInfo

	requestURL := fmt.Sprintf("%v://%v/%v/%v", githubAPI_SCHEME, githubAPI_URL, gp.Path, lastReleaseReq)

	body, err := retry.DoWithData(
		func() ([]byte, error) {
			res, err := gp.client.Get(requestURL)
			if err != nil {
				return nil, err
			}

			if res.StatusCode != http.StatusOK {
				log.Warnf("GitHub requests for repo %v answered with bad StatusCode: %v. Wait 10 min", gp.Path, res.StatusCode)
				return nil, errNo200
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}

			return body, nil
		},
		retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
			return time.Minute * 10
		}),
	)

	// err := retry.Do()

	if err != nil {
		fmt.Println(err)
		fmt.Println("LOL")
	}

	err = json.Unmarshal(body, &ri)
	if err != nil {
		return ri, err
	}

	return ri, nil
}

func (gp *GithubProvider) updateRelease(release ReleaseInfo) {
	gp.Release = release
}

func (gp *GithubProvider) getTitle() string {
	path := strings.Split(gp.Path, "/")

	return fmt.Sprintf("<b>%v</b>\n Release: <b>%v</b>\n", path[1], gp.Release.TagName)
}
