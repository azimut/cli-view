package twitter

import (
	"testing"
	"time"
)

const URL = `https://twitter.com/TwitterDev/status/1443269993676763138`

func TestTWFetch(t *testing.T) {
	_, err := Fetch(URL, "Twitter_View/0.1", time.Second*10)
	if err != nil {
		t.Fail()
	}
}
