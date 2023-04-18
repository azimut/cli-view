package fourchan

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

func parseUrl(rawUrl string) (int, string, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return -1, "", err
	}

	if uri.Host != "boards.4channel.org" {
		return -1, "", errors.New("Invalid hostname")
	}

	rawId := strings.Split(uri.Path, "/")[3]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return -1, "", err
	}
	board := strings.Split(uri.Path, "/")[1]

	return id, board, nil
}
