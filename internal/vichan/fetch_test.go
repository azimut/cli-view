package vichan

import (
	"fmt"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	url := "https://lainchan.org/sec/res/18084.html"
	thread, err := Fetch(url, "LainView/1.0", 100, 3, time.Second*5)
	if err != nil {
		t.Errorf("%v", err)
	}
	got := thread.op.id
	expected := 18084
	fmt.Println(len(thread.comments))
	fmt.Println(thread.comments)
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}
