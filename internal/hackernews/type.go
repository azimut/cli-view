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

func comment2comment(c gophernews.Comment, indent int) Comment {
	return Comment{
		id:     c.ID,
		msg:    c.Text,
		user:   c.By,
		kids:   c.Kids,
		date:   time.Unix(int64(c.Time), 0),
		indent: indent,
	}
}

func story2op(s gophernews.Story) Op {
	return Op{
		url:       s.URL,
		title:     s.Title,
		score:     s.Score,
		ncomments: len(s.Kids),
		user:      s.By,
		date:      time.Unix(int64(s.Time), 0),
	}
}
