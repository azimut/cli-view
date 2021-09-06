package hackernews

import "time"

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
	indent int
	user   string
	date   time.Time
}
