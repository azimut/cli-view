package fourchan

import (
	"fmt"
	"testing"
)

const URL = `https://boards.4channel.org/g/thread/92748329/sdg-stable-diffusion-general`

func TestFormatThread(t *testing.T) {
	thread, err := Fetch(URL)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(thread)
}
