package fourchan

import (
	"fmt"
	"testing"
)

const URL = `https://boards.4channel.org/g/thread/76759434/this-board-is-for-the-discussion-of-technology`

func TestFormatThread(t *testing.T) {
	thread, err := Fetch(URL, 100, 3)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(thread)
}
