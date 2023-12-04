package github

type ReleaseInfo struct {
	TagName string `json:"tag_name"`
	HtmlUrl string `json:"html_url"`
	Body    string `json:"body"`
}
