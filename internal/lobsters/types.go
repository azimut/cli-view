package lobsters

import "time"

type Thread struct {
	op           Op
	comments     []Comment
	CommentWidth int
	LeftPadding  int
	LineWidth    int
	ShowDate     bool
	ShowAuthor   bool
}

type Op struct {
	createdAt time.Time
	id        string
	message   string
	nComments int
	self      string
	score     int
	title     string
	thread    *Thread
	url       string
	username  string
}

type Comment struct {
	createdAt time.Time
	depth     int
	id        string
	message   string
	parent    string
	replies   []Comment
	score     int
	thread    *Thread
	username  string
}
