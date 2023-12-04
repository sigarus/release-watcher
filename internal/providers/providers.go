package providers

type Provider interface {
	WatchReleases() (title, description, link string, err error)
	GetName() string
}
