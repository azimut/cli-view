package hackernews

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestMakeComments(t *testing.T) {
	file, err := ioutil.ReadFile("/home/sendai/testfield/hn.html")
	if err != nil {
		t.Error(err)
	}
	r := bytes.NewReader(file)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Error(err)
	}
	comments := MakeComments(doc)
	if len(comments) != 19 {
		t.Errorf("expected 19 root comments got %d", len(comments))
	}
	if len(comments[0].Childs) != 4 {
		t.Errorf("expected 4 child comments for [0] got %d", len(comments[0].Childs))
	}
	if len(comments[0].Childs[0].Childs) != 1 {
		t.Errorf("expected 1 child comment for .Childs[0].Childs got %d", len(comments[0].Childs[0].Childs))
	}
}
