package fourchan

import "testing"

var testUrls = []string{
	"https://boards.4channel.org/g/thread/92730756/",
	"https://boards.4channel.org/g/thread/92730756/my-star64-arrived",
	"https://boards.4channel.org/g/thread/92730756/#p92721670",
}

func TestGetThreadId(t *testing.T) {
	for _, url := range testUrls {
		threadId, board, err := parseUrl(url)
		if err != nil || threadId != 92730756 || board != "g" {
			t.Fail()
		}
	}
}
