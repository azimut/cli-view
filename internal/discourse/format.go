package discourse

import (
	"fmt"

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
		format.FormatHtml2Text(o.message, o.thread.Width, o.thread.LeftPadding),
	)
	ret += fmt.Sprintf("%s\n\n\n", humanize.Time(o.createdAt))
	return
}

func (c Comment) String() (ret string) {
	ret += fmt.Sprintf(
		"%s\n",
		format.FormatHtml2Text(c.message, c.thread.Width, c.thread.LeftPadding*c.depth+1),
	)
	ret += fmt.Sprintf(">> %s %s\n", humanize.Time(c.createdAt), c.author)
	return
}
