package hackernews

import (
	"fmt"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/PuerkitoBio/goquery"
	"github.com/dustin/go-humanize"
	"jaytaylor.com/html2text"
)

func PrintDoc(doc *goquery.Document) {
	fmt.Println(NewHeader(doc))
	printChilds(NewComments(doc))
}

func printChilds(c []*Comment) {
	for _, value := range c {
		fmt.Println(value)
		printChilds(value.Childs)
	}
}

func (o *Op) String() (ret string) {
	ret += "title: " + o.title + "\n"
	ret += "url: " + o.url + "\n"
	ret += fmt.Sprintf("%s(%d) - %s\n", o.user, o.score, humanize.Time(o.date))
	return
}

func (c *Comment) String() (ret string) {
	indent := c.indent * 5
	msg, err := html2text.FromString(c.msg)
	if err != nil {
		panic(err)
	}
	wrapped, _ := text.WrapLeftPadded(msg, 120, indent+1)
	ret += "\n" + wrapped + "\n"
	arrow := ">>> "
	if c.indent > 0 {
		arrow = ">> "
	}
	ret += strings.Repeat(" ", indent) + arrow + c.user + " - " + humanize.Time(c.date)
	return
}
