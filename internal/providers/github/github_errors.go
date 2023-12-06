package github

import "errors"

var (
	errNo200     = errors.New("status code for request Gitlab API not 200")
	errRateLimit = errors.New("rate Limit error from GitHub")
)
