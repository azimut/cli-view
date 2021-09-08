package hackernews

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/azimut/cli-view/internal/fetch"
)

func Fetch(url, ua string, timeout time.Duration) (doc *goquery.Document, err error) {
	url, err = effectiveUrl(url)
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
