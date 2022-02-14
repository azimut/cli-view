package hackernews

import (
	"errors"
	"net/url"
)

func effectiveUrl(rawUrl string) (string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	if uri.Host != "news.ycombinator.com" {
		return "", errors.New("invalid hostname")
	}
	if uri.Path != "/item" {
		return "", errors.New("invalid path")
	}
	uri.Scheme = "https"
	return uri.String(), nil
}
