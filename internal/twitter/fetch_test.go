package twitter

import (
	"io/ioutil"
	"testing"
	"time"
)

const URL = `https://twitter.com/TwitterDev/status/1443269993676763138`

var TweetFiles = []string{
	"tweet-image-AND-quote.json",
	"tweet.json",
	"tweet-mult-image.json",
	"tweet-link.json",
	"tweet-quote.json",
	"tweet-url-with-image.json",
}

func TestTWFetch(t *testing.T) {
	_, err := Fetch(URL, "Twitter_View/0.1", time.Second*10)
	if err != nil {
		t.Fail()
	}
}

func TestUnmarshall(t *testing.T) {
	for _, tweetFile := range TweetFiles {
		bytes, err := ioutil.ReadFile("../../testdata/" + tweetFile)
		if err != nil {
			t.Fail()
			break
		}
		_, err = toEmbedded(string(bytes))
		if err != nil {
			t.Fail()
			break
		}
	}
}
