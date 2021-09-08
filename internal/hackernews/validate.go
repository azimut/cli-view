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
	return uri.String(), nil
}
