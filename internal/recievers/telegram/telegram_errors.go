package telegram

import "errors"

var (
	no200Err = errors.New("status code for request Telegram API not 200")
)
