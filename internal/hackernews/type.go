package hackernews

import (
	"time"

	"github.com/caser/gophernews"
)

type Op struct {
	url       string
	title     string
	score     int
	ncomments int
	user      string
	date      time.Time
	selfUrl   string
}

type Comment struct {
	id     int
	msg    string
	Childs []*Comment
	kids   []int
	indent int
	user   string
	date   time.Time
}

func newOp(story *gophernews.Story, selfUrl string) Op {
	return Op{
		date:      unix2time(story.Time),
		ncomments: len(story.Kids), // ?
		score:     story.Score,
		selfUrl:   selfUrl,
		title:     story.Title,
		url:       story.URL,
		user:      story.By,
	}
}

func unix2time(t int) time.Time {
	return time.Unix(int64(t), 0)
}
