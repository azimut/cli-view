package vichan

import (
	"fmt"
	"path"
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
	for _, attachment := range op.attachments {
		ret += fmt.Sprintf(
			" file: %s \"%s\"\n",
			path.Dir(path.Dir(op.thread.url))+"/src/"+attachment.newFilename,
			attachment.oldFilename,
		)
	}
	if op.message != "" {
		ret += "\n" + formatText(op.message, int(op.thread.width), int(op.thread.leftPadding))
		ret += "\n\n"
	}
	ret += " " + humanize.Time(op.createdAt)
	if op.thread.showAuthor {
		ret += " by " + op.author
	}
	ret += "\n\n\n"
	return
}

func (comment Comment) String() (ret string) {
	if comment.message != "" {
		ret += formatText(
			comment.message,
			int(comment.thread.width),
			comment.depth*int(comment.thread.leftPadding)+1,
		)
		ret += "\n"
	}

	ret += strings.Repeat(" ", comment.depth*3)
	ret += ">>"
	if comment.thread.showDate {
		ret += " " + humanize.Time(comment.createdAt)
	}
	if comment.thread.showAuthor {
		ret += " " + comment.author
	}
	if comment.thread.showId {
		ret += " " + fmt.Sprintf("%d", comment.id)
	}
	ret += "\n"

	for _, attachment := range comment.attachments {
		ret += strings.Repeat(" ", comment.depth*3)
		ret += fmt.Sprintf(
			">> %s \"%s\"\n",
			path.Dir(path.Dir(comment.thread.url))+"/src/"+attachment.newFilename,
			attachment.oldFilename,
		)
	}
	ret += "\n"

	for _, reply := range comment.replies {
		ret += fmt.Sprint(reply)
	}
	return
}
