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

func (thread Thread) String() (ret string) {
	ret += fmt.Sprintln(thread.op)
	for _, comment := range thread.comments {
		ret += fmt.Sprintln(comment)
	}
	return
}

func (o Op) String() (ret string) {
	ret += "\ntitle: " + o.title + "\n"
	if o.url != "" {
		ret += "  url: " + o.url + "\n"
		ret += " past: " + pastLink(o.title)
	}
	ret += fmt.Sprintf(" self: %s\n", o.selfUrl)
	if o.text != "" {
		ret += fmt.Sprintf("\n%s\n", fixupComment(o.text, 3, o.thread.LineWidth))
	}
	ret += fmt.Sprintf(
		"\n%s(%d) - %s - %d Comments\n",
		format.AuthorStyle.Render(o.user),
		o.score,
		humanize.Time(o.date),
		o.ncomments,
	)
	return
}

func (c Comment) String() (ret string) {
	leftPadding := c.thread.LeftPadding * c.depth
	rightPadding := 2
	lineWidth := format.Min(c.thread.LineWidth, leftPadding+c.thread.CommentWidth+1) - rightPadding
	ret += "\n" + fixupComment(c.msg, leftPadding+1, lineWidth) + "\n"

	arrow := ">> "
	if c.depth > 0 {
		arrow = ">> "
	}

	author := c.user
	if c.user == c.thread.op.user {
		author = format.AuthorStyle.Render(c.user)
	}

	if c.thread.ShowDate {
		ret += strings.Repeat(" ", leftPadding) + arrow + author + " - " + humanize.Time(c.date)
	} else {
		ret += strings.Repeat(" ", leftPadding) + arrow + author
	}

	ret += "\n"
	return
}

func fixupComment(html string, leftPad int, width int) string {
	plainText, err := html2text.FromString(
		html,
		html2text.Options{OmitLinks: false, PrettyTables: true, CitationStyleLinks: true},
	)
	if err != nil {
		panic(err)
	}
	wrapped, _ := text.WrapLeftPadded(format.GreenTextIt(plainText), width, leftPad)
	return wrapped
}

func pastLink(title string) string {
	return fmt.Sprintf(
		"https://hn.algolia.com/?query=%s&sort=byDate\n",
		url.QueryEscape(title),
	)
}
