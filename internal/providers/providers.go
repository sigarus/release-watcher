package providers

type Provider interface {
	WatchReleases() (name, release, description, link string, err error)
	GetName() string
}
