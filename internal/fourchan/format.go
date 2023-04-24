package fourchan

import (
	"fmt"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/azimut/cli-view/internal/format"
	"github.com/dustin/go-humanize"
)

func (thread Thread) String() (ret string) {
	ret += fmt.Sprint(thread.op)
	for _, post := range thread.posts {
		ret += fmt.Sprint(post)
		ret += "\n"
	}
	return
}

func (op Op) String() (ret string) {
	url := fmt.Sprintf("https://boards.4channel.org/%s/thread/%d/", op.board, op.id)
	if op.subject != "" {
		ret += fmt.Sprintf("title: %s\n", op.subject)
	}
	ret += fmt.Sprintf(" self: %s\n", url)
	if op.attachment.url != "" {
		ret += fmt.Sprintf("image: %s (%s)\n", op.attachment.url, op.attachment.filename)
	}
	ret += "\n"
	// TODO: better parser to handle links..etc..
	if op.comment != "" {
		comment, _ := text.WrapLeftPadded(
			format.GreenTextIt(op.comment),
			int(op.thread.width),
			int(op.thread.leftPadding),
		)
		ret += comment + "\n"
	}
	ret += "\n"
	ret += fmt.Sprintf("%s\n\n\n\n", humanize.Time(op.created))
	return
}

func (post Post) String() (ret string) {
	if post.comment != "" {
		comment, _ := text.WrapLeftPadded(
			format.GreenTextIt(post.comment),
			int(post.thread.width),
			post.depth*int(post.thread.leftPadding)+1,
		)
		ret += comment + "\n"
	}

	ret += strings.Repeat(" ", post.depth*3)
	if post.attachment.filename == "" {
		ret += fmt.Sprintf(">> %-13s", humanize.Time(post.created))
	} else {
		ret += fmt.Sprintf(">> %-13s | %s (%s)",
			humanize.Time(post.created),
			post.attachment.url,
			post.attachment.filename,
		)
	}
	ret += "\n\n"
	for _, reply := range post.replies {
		ret += fmt.Sprint(reply)
	}
	return
}
