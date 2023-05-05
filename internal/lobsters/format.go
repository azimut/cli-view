package lobsters

import (
	"fmt"
	"strings"

	"github.com/azimut/cli-view/internal/format"
	"github.com/dustin/go-humanize"
)

func (t Thread) String() (ret string) {
	ret += t.op.String()
	for _, comment := range t.comments {
		ret += comment.String()
		ret += "\n"
	}
	return
}

func (o Op) String() (ret string) {
	ret += fmt.Sprintf("title: %s\n", o.title)
	ret += fmt.Sprintf(" self: %s\n", o.self)
	if o.url != "" {
		ret += fmt.Sprintf("  url: %s\n", o.url)
	}
	ret += fmt.Sprintf(
		"\n%s\n\n",
		format.FormatHtml2Text(o.message, o.thread.LineWidth, o.thread.LeftPadding),
	)
	ret += fmt.Sprintf(
		"%s  - %s \n\n\n",
		format.AuthorStyle.Render(o.username),
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
		if c.username == c.thread.op.username {
			ret += format.AuthorStyle.Render(c.username)
		} else {
			ret += c.username
		}
	}
	if c.thread.ShowDate {
		ret += " " + humanize.Time(c.createdAt)
	}
	ret += "\n\n"
	// for _, reply := range c.replies {
	// 	ret += fmt.Sprint(reply)
	// }
	return
}
