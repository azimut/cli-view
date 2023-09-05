package discourse

import (
	"fmt"
	"regexp"
)

var urlRegex = regexp.MustCompile(`http[s]?://[^/]+/t/[^/]+/[0-9]+`)

func effectiveUrl(rawUrl string) (string, error) {
	url := urlRegex.FindString(rawUrl)
	if url == "" {
		return "", fmt.Errorf("invalid url: %s", rawUrl)
	}
	return url + ".json", nil
}
