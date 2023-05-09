package discourse

import (
	"fmt"
	"strings"

	"github.com/azimut/cli-view/internal/format"
	"github.com/dustin/go-humanize"
)

func (t Thread) String() (ret string) {
	ret += fmt.Sprint(t.op)
	for _, comment := range t.comments {
		ret += fmt.Sprint(comment)
		ret += "\n"
	}
	return
}

func (o Op) String() (ret string) {
	ret += fmt.Sprintf("title: %s\n", o.title)
	ret += fmt.Sprintf(" self: %s\n", o.thread.url)
	ret += fmt.Sprintf(
		"\n%s\n\n",
		format.FormatHtml2Text(
			o.message,
			o.thread.LineWidth-o.thread.LeftPadding,
			o.thread.LeftPadding,
		),
	)
	ret += fmt.Sprintf(
		"%s  - %s \n\n\n",
		format.AuthorStyle.Render(o.author),
		humanize.Time(o.createdAt),
	)
	return
}

func (c Comment) String() (ret string) {
	leftPadding := c.thread.LeftPadding * c.depth
	rightPadding := 2
	extraLeft := 1
	lineWidth := format.Min(
		c.thread.LineWidth,
		leftPadding+c.thread.CommentWidth+extraLeft,
	) - rightPadding
	ret += fmt.Sprintf(
		"%s\n",
		format.FormatHtml2Text(c.message, lineWidth, leftPadding+extraLeft),
	)
	ret += strings.Repeat(" ", leftPadding)
	ret += ">>"
	if c.thread.ShowAuthor {
		ret += " "
		if c.author == c.thread.op.author {
			ret += format.AuthorStyle.Render(c.author)
		} else {
			ret += c.author
		}
	}
	if c.thread.ShowDate {
		ret += " " + humanize.Time(c.createdAt)
	}
	ret += "\n\n"
	for _, reply := range c.replies {
		ret += fmt.Sprint(reply)
	}
	return
}
