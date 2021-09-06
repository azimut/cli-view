package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"github.com/azimut/cli-view/internal/hackernews"
)

func main() {
	file, err := ioutil.ReadFile("/home/sendai/testfield/hn.html")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(file)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		panic(err)
	}
	comments := hackernews.MakeComments(doc)
	fmt.Println(comments[0].Childs[0])
}
