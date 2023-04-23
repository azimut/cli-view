package vichan

import (
	"errors"
	"net/url"
	"strings"
)

func parseUrl(rawUrl string) (string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(uri.Path, ".html") {
		return "", errors.New("invalid url, path does not end with .html")
	}

	uri.RawQuery = ""

	return strings.TrimSuffix(uri.String(), ".html") + ".json", nil
}
