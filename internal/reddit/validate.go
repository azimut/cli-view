package reddit

import (
	"fmt"
	"net/url"
)

func effectiveUrl(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	if u.Host == "old.reddit.com" || u.Host == "reddit.com" {
		u.Host = "www.reddit.com"
	}
	if u.Host != "www.reddit.com" {
		return "", fmt.Errorf("not supported host: %s", u.Host)
	}
	return uri + ".json", nil
}
