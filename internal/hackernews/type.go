package hackernews

import (
	"time"

	"github.com/caser/gophernews"
)

type Thread struct {
	comments     []Comment
	op           Op
	CommentWidth int
	LineWidth    int
	LeftPadding  int
	ShowDate     bool
}

type Op struct {
	date      time.Time
	kids      []int
	ncomments int
	score     int
	selfUrl   string
	text      string
	title     string
	url       string
	user      string
	thread    *Thread
}

type Comment struct {
	Childs []*Comment
	date   time.Time
	id     int
	depth  int
	kids   []int
	msg    string
	user   string
	thread *Thread
}

func newComment(comment gophernews.Comment, thread *Thread) Comment {
	return Comment{
		date:   unix2time(comment.Time),
		id:     comment.ID,
		kids:   comment.Kids,
		msg:    comment.Text,
		thread: thread,
		user:   comment.By,
	}
}

func unix2time(t int) time.Time {
	return time.Unix(int64(t), 0)
}
