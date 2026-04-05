package utils

import (
	"errors"
	"net/url"
	"strings"
)

var ErrInvalidScheme = errors.New("invalid url scheme: must be http or https")

func ValidateURL(urlStr string) error {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}

	scheme := strings.ToLower(parsedURL.Scheme)
	if scheme != "http" && scheme != "https" {
		return ErrInvalidScheme
	}

	return nil
}
