package discourse

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

func effectiveUrl(rawUrl string) (string, error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	paths := strings.Split(url.Path, "/")
	if len(paths) != 4 {
		return "", errors.New("invalid url, does not enough paths")
	}

	if paths[1] != "t" {
		return "", errors.New("invalid url, does not have a `/t/` path")
	}

	rawId := paths[3]
	if _, err = strconv.Atoi(rawId); err != nil {
		return "", errors.New("invalid url, does not end with a number")
	}

	return url.String() + ".json", nil
}
