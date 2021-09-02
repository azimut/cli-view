package twitter

import (
	"fmt"
	"regexp"

	"jaytaylor.com/html2text"
)

func enrich(html string) (md string) {
	md, err := html2text.FromString(html)
	if err != nil {
		panic(err)
	}
	return
}

func extract(html string) (p string) {
	r, err := regexp.Compile("<p.*>(.*)</p>")
	if err != nil {
		panic(err)
	}
	return r.FindStringSubmatch(html)[0]
}

func Format(t *Embedded) (formatted string) {
	formatted += fmt.Sprintf("URL: %s\n", t.Url)
	formatted += enrich(extract(t.Html))
	formatted += fmt.Sprintf("> %s\n", t.AuthorName)
	return
}
