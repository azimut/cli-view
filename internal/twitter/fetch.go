package twitter

import (
	"encoding/json"
	"time"

	"github.com/azimut/cli-view/internal/fetch"
)

func Fetch(url, ua string, timeout time.Duration) (tweet *Embedded, err error) {
	url, err = EffectiveUrl(url)
	if err != nil {
		return nil, err
	}
	res, err := fetch.Fetch(url, ua, timeout)
	if err != nil {
		return nil, err
	}
	b := []byte(res)
	if err = json.Unmarshal(b, &tweet); err != nil {
		return nil, err
	}
	return
}
