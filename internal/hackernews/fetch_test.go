package hackernews

import (
	"testing"
	"time"
)

const URL = `https://news.ycombinator.com/item?id=3078128`

func TestHNFetch(t *testing.T) {
	_, _, err := Fetch(URL, time.Second*10, 10, 2)
	if err != nil {
		t.Fail()
	}
}
