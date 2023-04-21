package reddit

import (
	"testing"
	"time"
)

const URL = `https://old.reddit.com/r/gamedev/comments/horm1a/ive_started_making_a_teaching_aid_for_people/`

func TestRTFetch(t *testing.T) {
	_, err := Fetch(URL, "Reddit_Cli/0.1", 100, 3, time.Second*10)
	if err != nil {
		t.Fail()
	}
}
