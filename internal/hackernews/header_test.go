package hackernews

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestNewHeader(t *testing.T) {
	file, err := ioutil.ReadFile("../../testdata/hn.html")
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(file)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatal(err)
	}
	header := NewHeader(doc)
	if header.url != "https://www.newyorker.com/news/news-desk/the-red-warning-light-on-richard-bransons-space-flight" {
		t.Errorf("invalid URL")
	}
	if header.score != 312 {
		t.Errorf("invalid score, expected 312")
	}
	if header.user != "zlsa" {
		t.Errorf("invalid user, expected \"zlsa\"")
	}
}
