package fourchan

import "time"

type Thread struct {
	closed bool
	id     int
	op     Op
	posts  []Post
}

type Op struct {
	attachment Attachment
	board      string
	comment    string
	created    time.Time
	id         int
	subject    string
}

type Post struct {
	attachment Attachment
	comment    string
	created    time.Time
	depth      int
	replies    []Post
}

type Attachment struct {
	filename string
	id       string
	url      string
}
