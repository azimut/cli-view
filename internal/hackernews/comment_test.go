package hackernews

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestMakeComments(t *testing.T) {
	file, err := ioutil.ReadFile("../../testdata/hn.html")
	if err != nil {
		t.Fatal(err)
	}
	r := bytes.NewReader(file)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatal(err)
	}
	comments := NewComments(doc)
	result := len(comments)
	if result != 19 {
		t.Errorf("expected 19 root comments got %d", result)
	}
	result = len(comments[0].Childs)
	if result != 4 {
		t.Errorf("expected 4 child comments for [0].Childs got %d", result)
	}
	result = len(comments[0].Childs[0].Childs)
	if result != 1 {
		t.Errorf("expected 1 child comments for [0].Childs[0].Childs got %d", result)
	}
}
