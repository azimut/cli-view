package hackernews

import (
	"time"
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
