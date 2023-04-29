package vichan

import (
	"fmt"
	"path"
	"strings"

	"github.com/azimut/cli-view/internal/format"

	"github.com/dustin/go-humanize"
)

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
		ret += "\n" + format.FormatHtml2Text(
			op.message,
			op.thread.LineWidth,
			op.thread.LeftPadding,
		)
		ret += "\n\n"
	}
	if op.thread.ShowAuthor {
		ret += format.AuthorStyle.Render(op.author)
	}
	ret += " " + humanize.Time(op.createdAt)
	ret += "\n\n\n"
	return
}

func (comment Comment) String() (ret string) {
	leftPadding := comment.thread.LeftPadding * comment.depth
	rightPadding := 2
	lineWidth := format.Min(
		comment.thread.LineWidth,
		leftPadding+comment.thread.CommentWidth+1,
	) - rightPadding
	if comment.message != "" {
		ret += format.FormatHtml2Text(
			comment.message,
			lineWidth,
			leftPadding+1,
		)
		ret += "\n"
	}

	ret += strings.Repeat(" ", comment.depth*3)
	ret += ">>"
	if comment.thread.ShowDate {
		ret += " " + humanize.Time(comment.createdAt)
	}
	if comment.thread.ShowAuthor {
		ret += " "
		if comment.author == comment.thread.op.author {
			ret += format.AuthorStyle.Render(comment.author)
		} else {
			ret += comment.author
		}
	}
	if comment.thread.ShowId {
		ret += " " + fmt.Sprintf("%d", comment.id)
	}
	ret += "\n"

	for _, attachment := range comment.attachments {
		ret += strings.Repeat(" ", leftPadding)
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
