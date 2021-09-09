package twitter

import (
	"fmt"
	"strings"
	"time"

	text "github.com/MichaelMure/go-term-text"
	"github.com/PuerkitoBio/goquery"
	"github.com/dustin/go-humanize"
	"jaytaylor.com/html2text"
)

func externalName(external string) string {
	if strings.Contains(external, "pic.twitter.com") {
		return "image: " + external + "\n"
	} else if strings.Contains(external, "t.co") {
		return "external: " + external + "\n"
	} else if external != "" {
		panic("unknown not external URL passed: " + external)
	}
	return ""
}

func Format(t *Embedded) (formatted string) {
	msg, links := paragraph(t.Html)
	formatted += fmt.Sprintf("URL: %s\n", t.Url)
	for _, link := range links {
		formatted += externalName(link)
	}
	formatted += "\n" + fitInScreen(plaintext(msg))
	formatted += "\n\n" + date(t.Html)
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

func date(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}
	doc.Find("p").Remove() // Msg
	rawdate, err := doc.Find("a").Html()
	if err != nil {
		panic(err)
	}
	date, err := time.Parse("January 2, 2006", rawdate)
	if err != nil {
		panic(err)
	}
	return humanize.Time(date)
}

func paragraph(html string) (string, []string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}
	removeAllHashtagLinks(doc.Find("p a"))
	removeAllMentionLinks(doc.Find("p a"))
	external := popExternalLinks(doc.Find("p a"))
	msg, err := doc.Find("p").Html()
	if err != nil {
		panic(err)
	}
	return msg, external
}

func popExternalLinks(sel *goquery.Selection) (ret []string) {
	sel.Each(func(i int, s *goquery.Selection) {
		link := s.Text()
		if strings.HasPrefix(link, "pic.twitter.com") {
			link = "https://" + link
		}
		ret = append(ret, link)
		s.Remove()
	})
	return ret
}

func removeAllMentionLinks(sel *goquery.Selection) {
	sel.Each(func(i int, s *goquery.Selection) {
		mention := s.Text()
		if strings.HasPrefix(mention, "@") {
			s.ReplaceWithHtml(mention)
		}
	})
}

func removeAllHashtagLinks(sel *goquery.Selection) {
	sel.Each(func(i int, s *goquery.Selection) {
		hashtag := s.Text()
		if strings.HasPrefix(hashtag, "#") {
			s.ReplaceWithHtml(hashtag)
		}
	})
}
