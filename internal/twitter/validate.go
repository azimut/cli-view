package twitter

import (
	"errors"
	"net/url"
	"regexp"
)

func EffectiveUrl(rawUrl string) (string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	if uri.Host == "m.twitter.com" || uri.Host == "mobile.twitter.com" || uri.Host == "t.co" {
		uri.Host = "twitter.com"
	}
	if uri.Host != "twitter.com" {
		return "", errors.New("invalid hostname")
	}
	uri.RawQuery = ""
	uri.Scheme = "https"
	reTrimmer := regexp.MustCompile("/photo/[0-9]+$|/video/[0-9]+$|/retweets/with_comments$")
	uri.Path = reTrimmer.ReplaceAllString(uri.Path, "")
	return "https://publish.twitter.com/oembed?url=" + uri.String(), nil
}
