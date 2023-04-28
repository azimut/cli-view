package discourse

import (
	"time"

	"github.com/azimut/cli-view/internal/fetch"
)

func Fetch(rawUrl, userAgent string, timeout time.Duration) (*Thread, error) {
	url, err := effectiveUrl(rawUrl)
	if err != nil {
		return nil, err
	}

	rawJson, err := fetch.Fetch(url, userAgent, timeout)
	if err != nil {
		return nil, err
	}

	thread, err := toThread(rawJson, rawUrl)
	if err != nil {
		return nil, err
	}

	return thread, nil
}
