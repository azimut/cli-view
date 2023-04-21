package hackernews

import (
	"fmt"
	"net/url"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/azimut/cli-view/internal/format"
	"github.com/dustin/go-humanize"
	"github.com/jaytaylor/html2text"
)

const SPACES_PER_INDENT = 5

var max_width int
var useDate bool

func Format(width int, useDateArg bool, op Op, comments []Comment) {
	max_width = width
	useDate = useDateArg
	fmt.Println(op)
	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func printChilds(c []*Comment) {
	for _, value := range c {
		fmt.Println(value)
		printChilds(value.Childs)
	}
}

func pastLink(title string) string {
	return fmt.Sprintf(
		"https://hn.algolia.com/?query=%s&sort=byDate\n",
		url.QueryEscape(title),
	)
}

func (o Op) String() (ret string) {
	ret += "\ntitle: " + o.title + "\n"
	if o.url != "" {
		ret += "  url: " + o.url + "\n"
		ret += " past: " + pastLink(o.title)
	}
	ret += fmt.Sprintf(" self: %s\n", o.selfUrl)
	ret += fmt.Sprintf(
		"\n%s(%d) - %s - %d Comments\n",
		o.user,
		o.score,
		humanize.Time(o.date),
		o.ncomments,
	)
	return
}

func (c Comment) String() (ret string) {
	indent := c.indent * SPACES_PER_INDENT
	msg, err := html2text.FromString(
		c.msg,
		html2text.Options{OmitLinks: false, PrettyTables: true, CitationStyleLinks: true},
	)

	if err != nil {
		panic(err)
	}
	wrapped, _ := text.WrapLeftPadded(format.GreenTextIt(msg), max_width, indent+1)
	ret += "\n" + wrapped + "\n"
	arrow := ">> "
	if c.indent > 0 {
		arrow = ">> "
	}
	if useDate {
		ret += strings.Repeat(" ", indent) + arrow + c.user + " - " + humanize.Time(c.date)
	} else {
		ret += strings.Repeat(" ", indent) + arrow + c.user
	}
	ret += "\n"
	return
}
