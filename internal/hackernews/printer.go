package hackernews

import (
	"fmt"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/dustin/go-humanize"
	"github.com/jaytaylor/html2text"
)

const SPACES_PER_INDENT = 5

var max_width int

func Format(width int, op *Op, comments *[]Comment) {
	max_width = width
	fmt.Println(op)
	for _, comment := range *comments {
		fmt.Println(&comment)
	}
}

func printChilds(c []*Comment) {
	for _, value := range c {
		fmt.Println(value)
		printChilds(value.Childs)
	}
}

func (o *Op) String() (ret string) {
	ret += "title: " + o.title + "\n"
	if o.url != "" {
		ret += "url: " + o.url + "\n"
	}
	ret += "self: " + o.selfUrl + "\n"
	ret += fmt.Sprintf(
		"%s(%d) - %s - %d Comments\n",
		o.user,
		o.score,
		humanize.Time(o.date),
		o.ncomments,
	)
	return
}

func (c *Comment) String() (ret string) {
	indent := c.indent * SPACES_PER_INDENT
	msg, err := html2text.FromString(
		c.msg,
		html2text.Options{OmitLinks: false, PrettyTables: true, CitationStyleLinks: true},
	)
	if err != nil {
		panic(err)
	}
	wrapped, _ := text.WrapLeftPadded(msg, max_width, indent+1)
	ret += "\n" + wrapped + "\n"
	arrow := ">>> "
	if c.indent > 0 {
		arrow = ">> "
	}
	ret += strings.Repeat(" ", indent) + arrow + c.user + " - " + humanize.Time(c.date)
	return
}
