package twitter

import (
	"fmt"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/PuerkitoBio/goquery"
	"jaytaylor.com/html2text"
)

func Format(t *Embedded) (formatted string) {
	formatted += fmt.Sprintf("URL: %s\n", t.Url)
	formatted += "\n" + fitInScreen(plaintext(paragraph(t.Html)))
	return
}

func fitInScreen(s string) string {
	s, _ = text.WrapLeftPadded(s, 100, 3)
	return s
}

func plaintext(html string) string {
	md, err := html2text.FromString(html)
	if err != nil {
		panic(err)
	}
	return md
}

func paragraph(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}
	doc.Find("p a").Remove()
	msg, err := doc.Find("p").Html()
	if err != nil {
		panic(err)
	}
	return msg
}
