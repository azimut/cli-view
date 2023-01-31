package twitter

import (
	"fmt"
	"strings"
	"time"

	text "github.com/MichaelMure/go-term-text"
	"github.com/PuerkitoBio/goquery"
	"github.com/azimut/cli-view/internal/fetch"
	"github.com/dustin/go-humanize"
	"github.com/jaytaylor/html2text"
)

func Format(t *Embedded) (formatted string) {
	msg, links := paragraph(t.Html)
	formatted += fmt.Sprintf("URL: %s\n", t.Url)
	for _, link := range links {
		formatted += formatLink(link)
	}
	formatted += "\n" + fitInScreen(plaintext(msg))
	formatted += "\n\n" + date(t.Html)
	return
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

func formatLink(link string) string {
	if strings.Contains(link, "pic.twitter.com") {
		return "<< HAS MEDIA >>\n"
	} else if strings.Contains(link, "t.co") {
		finalUrl, err := fetch.UrlLocation(link, "Mozilla", time.Second*10) // TODO: use arguments
		if err != nil {
			return "external: " + link + "\n"
		}
		return "external: " + finalUrl + "\n"
	} else if link != "" {
		panic("unknown not external URL passed: " + link)
	}
	return ""
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

func popExternalLinks(sel *goquery.Selection) (ret []string) {
	sel.Each(func(i int, s *goquery.Selection) {
		link := s.Text()
		s.Remove()
		if strings.HasPrefix(link, "pic.twitter.com") {
			link = "https://" + link
		}
		ret = append(ret, link)
	})
	return
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
