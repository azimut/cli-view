package lobsters

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var ErrInvalidDomain = errors.New("invalid domain")
var ErrInvalidUrl = errors.New("invalid url shape")

func effectiveUrl(rawUrl string) (string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	if uri.Hostname() != "lobste.rs" {
		return "", ErrInvalidDomain
	}
	paths := strings.Split(uri.Path, "/")
	if len(paths) < 3 || paths[1] != "s" {
		return "", ErrInvalidUrl
	}
	return fmt.Sprintf("https://lobste.rs/s/%s.json", paths[2]), nil
}
