package hackernews

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func NewHeader(doc *goquery.Document) *Op {
	return &Op{
		ncomments: getNComments(doc),
		title:     getTitle(doc),
		score:     getScore(doc),
		user:      getUser(doc),
		date:      getDate(doc),
		url:       getUrl(doc),
	}
}

func getUser(doc *goquery.Document) string {
	return doc.Find("td.subtext a.hnuser").Text()
}

func getDate(doc *goquery.Document) time.Time {
	rawdate, exists := doc.Find("td.subtext span.age").Attr("title")
	if !exists {
		panic("could not find span.age title")
	}
	date, err := time.Parse("2006-01-02T15:04:05", rawdate)
	if err != nil {
		panic(err)
	}
	return date
}

func TrimToNum(r rune) bool {
	if n := r - '0'; n >= 0 && n <= 9 {
		return false
	}
	return true
}

func getNComments(doc *goquery.Document) int {
	rawn, err := doc.Find("td.subtext a").Last().Html()
	if err != nil {
		panic(err)
	}
	rawn = strings.TrimFunc(rawn, TrimToNum)
	n, err := strconv.Atoi(rawn)
	if err != nil {
		panic(err)
	}
	return n
}

func getTitle(doc *goquery.Document) string {
	anchor, exists := doc.Find("tr#pagespace").Attr("title")
	if !exists {
		panic("no title")
	}
	return anchor
}

func getUrl(doc *goquery.Document) string {
	url, exists := doc.Find("a.storylink").Attr("href")
	if !exists {
		panic("no story href")
	}
	return url
}

func getScore(doc *goquery.Document) int {
	rawscore := doc.Find("span.score").Text()
	rawscore = strings.TrimRight(rawscore, " points")
	score, err := strconv.Atoi(rawscore)
	if err != nil {
		panic(err)
	}
	return score
}
