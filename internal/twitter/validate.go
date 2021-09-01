package twitter

import (
	"errors"
	"net/url"
)

const prefix = "https://publish.twitter.com/oembed?url="

func EffectiveUrl(rawUrl string) (string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	if uri.Host == "m.twitter.com" || uri.Host == "mobile.twitter.com" {
		uri.Host = "twitter.com"
	}
	if uri.Host != "twitter.com" {
		return "", errors.New("invalid hostname")
	}
	return prefix + uri.String(), nil
}
