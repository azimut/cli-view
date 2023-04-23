package vichan

import (
	"fmt"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/azimut/cli-view/internal/format"

	"github.com/dustin/go-humanize"
	"github.com/jaytaylor/html2text"
)

func formatText(htmlText string, width, leftPadding int) string {
	plainText, err := html2text.FromString(htmlText, html2text.Options{})
	if err != nil {
		panic(err)
	}
	wrapped, _ := text.WrapLeftPadded(format.GreenTextIt(plainText), width, leftPadding)
	return wrapped
}

func (thread Thread) String() (ret string) {
	ret += "\n"
	ret += fmt.Sprint(thread.op)
	for _, comment := range thread.comments {
		ret += fmt.Sprint(comment)
		ret += "\n"
	}
	return
}

func (op Op) String() (ret string) {
	if op.title != "" {
		ret += fmt.Sprintf("title: %s\n", op.title)
	}
	ret += fmt.Sprintf(" self: %s\n", op.thread.url)
	if op.message != "" {
		ret += "\n" + formatText(op.message, int(op.thread.width), int(op.thread.leftPadding))
	}
	ret += fmt.Sprintf("\n\n%s\n\n\n", humanize.Time(op.createdAt))
	return
}

func (comment Comment) String() (ret string) {
	if comment.message != "" {
		ret += formatText(
			comment.message,
			int(comment.thread.width),
			comment.depth*int(comment.thread.leftPadding)+1,
		)
	}
	ret += "\n" + strings.Repeat(" ", comment.depth*3)
	ret += fmt.Sprintf(">> %s\n\n", humanize.Time(comment.createdAt))
	for _, reply := range comment.replies {
		ret += fmt.Sprint(reply)
	}
	return
}
