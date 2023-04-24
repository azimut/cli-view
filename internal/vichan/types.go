package vichan

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Thread struct {
	op          Op
	comments    []Comment
	leftPadding uint
	width       uint
	url         string
	showAuthor  bool
	showDate    bool
	showId      bool
}

type Attachment struct {
	oldFilename string
	newFilename string
	height      int
	sizeInBytes int
	width       int
}

type Op struct {
	attachments []Attachment
	author      string
	createdAt   time.Time
	id          int
	message     string
	thread      *Thread
	title       string
}

type Comment struct {
	attachments []Attachment
	author      string
	createdAt   time.Time
	depth       int
	id          int
	message     string
	parentId    int
	replies     []Comment
	subject     string
	thread      *Thread
}

var reReplies = regexp.MustCompile(
	`<a onclick="highlightReply\('([0-9]+)', event\);" href="/[^/]+/res/[0-9]+.html#[0-9]+">&gt;&gt;[0-9]+</a>([ ]*<br/>)?`,
)

func (comment Comment) explode() (comments []Comment, err error) {
	subMatches := reReplies.FindAllStringSubmatch(comment.message, -1)
	ids := make([]int, len(subMatches))
	for i, subMatch := range subMatches {
		id, err := strconv.Atoi(subMatch[1])
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	replies := reReplies.Split(comment.message, -1)

	// Single comment, no replies
	if len(ids) == 0 {
		comment.message = replies[0]
		return append(comments, comment), nil
	}

	// Part of it replies to OP
	if replies[0] != "" {
		newComment := comment
		newComment.message = replies[0]
		comments = append(comments, newComment)
	}

	// Remove the not reply part
	replies = replies[1:]

	// is replying to many with the same message, clean it up
	if allEmptyButLast(replies) {
		replies = replies[len(replies)-1:] // keep only last in slice
		ids = ids[:1]                      // keep only first one
	}

	// 1/1 comment/reply
	if len(ids) == len(replies) && !anyEmptyString(replies) {
		for i, reply := range replies {
			newComment := comment
			newComment.message = reply
			newComment.parentId = ids[i]
			comments = append(comments, newComment)
		}
		return comments, nil
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("------ Comment", len(replies), `"`+comment.subject+`"`)
	fmt.Println(strings.ReplaceAll(comment.message, "<br/>", "<br/>\n"))
	fmt.Println("------ Findings")
	fmt.Println("ids: ", ids)
	fmt.Println("------ Replies")
	for _, reply := range replies {
		fmt.Printf("'%s'\n", reply)
	}
	return comments, nil
}

// insert assumes is already exploted
func (thread *Thread) insert(comment Comment) {
	if comment.parentId == 0 || comment.parentId == thread.op.id {
		thread.comments = append(thread.comments, comment)
	} else {
		parentComment := thread.getComment(comment.parentId)
		comment.depth = parentComment.depth + 1
		parentComment.replies = append(parentComment.replies, comment)
	}
}

func (thread *Thread) getComment(id int) *Comment {
	for i := range thread.comments {
		if found := thread.comments[i].getComment(id); found != nil {
			return found
		}
	}
	return nil
}

func (comment *Comment) getComment(id int) *Comment {
	if comment.id == id {
		return comment
	}
	for i := range comment.replies {
		if found := comment.replies[i].getComment(id); found != nil {
			return found
		}
	}
	return nil
}

func anyEmptyString(xs []string) bool {
	for _, x := range xs {
		if x == "" {
			return true
		}
	}
	return false
}

func allEmptyButLast(replies []string) bool {
	if len(replies) == 0 {
		return false
	}
	for _, reply := range replies[:len(replies)-1] {
		if reply != "" {
			return false
		}
	}
	return replies[len(replies)-1] != ""
}
