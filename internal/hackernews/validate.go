package hackernews

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func effectiveUrl(rawUrl string) (string, int, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", 0, err
	}
	if uri.Host != "news.ycombinator.com" {
		return "", 0, fmt.Errorf("invalid Host: %s", uri.Host)
	}
	if uri.Path != "/item" {
		return "", 0, fmt.Errorf("invalid Path: %s", uri.Path)
	}
	params := strings.Split(uri.RawQuery, "=")
	if params[0] != "id" || len(params) != 2 {
		return "", 0, fmt.Errorf("invalid RawQuery: %s", params)
	}
	id, err := strconv.Atoi(params[1])
	if err != nil {
		return "", 0, err
	}
	uri.Scheme = "https"
	return uri.String(), id, nil
}
