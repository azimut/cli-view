package fetch

import (
	"testing"
	"time"
)

const URL = `https://www.google.com/robots.txt`

func TestFech(t *testing.T) {
	_, err := Fetch(URL, "Wget", time.Second*10)
	if err != nil {
		t.Fail()
	}
}
