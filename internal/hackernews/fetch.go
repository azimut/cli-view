package hackernews

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/azimut/cli-view/internal/fetch"
)

func EffectiveUrl(rawUrl string) (string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	if uri.Host != "news.ycombinator.com" {
		return "", errors.New("invalid hostname")
	}
	return uri.String(), nil
}

func Fetch(url, ua string, timeout time.Duration) (doc *goquery.Document, err error) {
	url, err = EffectiveUrl(url)
	if err != nil {
		return nil, err
	}
	res, err := fetch.Fetch(url, ua, timeout)
	if err != nil {
		return nil, err
	}
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		return nil, err
	}
	return
}
