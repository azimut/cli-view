package fourchan

import (
	"fmt"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/jaytaylor/html2text"
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
	ret += fmt.Sprintf("title: %s\n", op.subject)
	ret += fmt.Sprintf(" self: %s\n", url)
	if op.attachment.filename != "" {
		ret += fmt.Sprintf("image: %s (%s)\n", op.attachment.url, op.attachment.filename)
	}
	ret += "\n"
	// TODO: better parser to handle links..etc..
	if op.comment != "" {
		comment, err := html2text.FromString(op.comment, html2text.Options{})
		if err != nil {
			panic(err)
		}
		ret += comment + "\n"
	}
	ret += "\n"
	ret += fmt.Sprintf(">>> %s\n\n", humanize.Time(op.created))
	return
}

func (post Post) String() (ret string) {
	if post.comment != "" {
		ret += post.comment + "\n"
	}

	if post.attachment.filename == "" {
		ret += fmt.Sprintf(">> %13s", humanize.Time(post.created))
	} else {
		if strings.HasSuffix(post.attachment.url, post.attachment.filename) {
			ret += fmt.Sprintf(">> %13s | %s | %s",
				humanize.Time(post.created),
				post.attachment.url,
				post.attachment.filename,
			)
		} else {
			ret += fmt.Sprintf(">> %13s | %s",
				humanize.Time(post.created),
				post.attachment.url,
			)
		}
	}
	ret += "\n\n"
	return
}
