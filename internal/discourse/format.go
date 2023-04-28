package discourse

import (
	"fmt"
	"strings"

	"github.com/azimut/cli-view/internal/format"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

var authorStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("8")).
	Foreground(lipgloss.Color("0"))

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
		format.FormatHtml2Text(o.message, o.thread.Width, o.thread.LeftPadding),
	)
	ret += fmt.Sprintf("%s  - %s \n\n\n", authorStyle.Render(o.author), humanize.Time(o.createdAt))
	return
}

func (c Comment) String() (ret string) {
	ret += fmt.Sprintf(
		"%s\n",
		format.FormatHtml2Text(c.message, c.thread.Width, c.thread.LeftPadding*c.depth+1),
	)
	ret += strings.Repeat(" ", c.depth*c.thread.LeftPadding)
	ret += ">>"
	if c.thread.ShowAuthor {
		ret += " "
		if c.author == c.thread.op.author {
			ret += authorStyle.Render(c.author)
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
