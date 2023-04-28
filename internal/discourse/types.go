package discourse

import "time"

type Thread struct {
	comments    []Comment
	op          Op
	id          int
	url         string
	Width       int
	LeftPadding int
}

type Op struct {
	author    string
	createdAt time.Time
	id        int
	message   string
	thread    *Thread
	title     string
}

type Comment struct {
	author    string
	createdAt time.Time
	depth     int
	id        int
	message   string
	replies   []Comment
	thread    *Thread
}
