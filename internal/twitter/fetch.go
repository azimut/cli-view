package twitter

import (
	"encoding/json"
	"time"

	"github.com/azimut/cli-view/internal/fetch"
)

func Fetch(tweetUrl, ua string, timeout time.Duration) (*Embedded, error) {
	embedUrl, err := EffectiveUrl(tweetUrl)
	if err != nil {
		return nil, err
	}
	rawJson, err := fetch.Fetch(embedUrl, ua, timeout)
	if err != nil {
		return nil, err
	}
	tweet, err := toEmbedded(rawJson)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func toEmbedded(rawJson string) (tweet *Embedded, err error) {
	b := []byte(rawJson)
	if err = json.Unmarshal(b, &tweet); err != nil {
		return nil, err
	}
	return
}
