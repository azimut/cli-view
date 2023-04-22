package hackernews

import (
	"time"

	"github.com/caser/gophernews"
)

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
}

type Comment struct {
	Childs []*Comment
	date   time.Time
	id     int
	indent int
	kids   []int
	msg    string
	user   string
}

func newOp(story gophernews.Item, selfUrl string) Op {
	return Op{
		date:      unix2time(story.Time()),
		kids:      story.Kids(),
		ncomments: len(story.Kids()), // TODO: this only gets the direct replies
		score:     story.Score(),
		selfUrl:   selfUrl,
		text:      story.Text(),
		title:     story.Title(),
		url:       story.URL(),
		user:      story.By(),
	}
}

func newComment(comment gophernews.Comment) Comment {
	return Comment{
		id:   comment.ID,
		msg:  comment.Text,
		user: comment.By,
		kids: comment.Kids,
		date: unix2time(comment.Time),
	}
}

func unix2time(t int) time.Time {
	return time.Unix(int64(t), 0)
}
