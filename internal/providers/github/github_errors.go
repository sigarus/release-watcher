package github

import "errors"

var (
	errNo200 = errors.New("status code for request Gitlab API not 200")
)
