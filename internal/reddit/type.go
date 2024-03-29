package reddit

import "github.com/tidwall/gjson"

type Thread struct {
	comments     []Comment
	op           Op
	LineWidth    int
	CommentWidth int
	LeftPadding  int
}

type Op struct {
	author     string
	createdUTC int64
	nComments  int64
	self       string
	selftext   string
	thread     *Thread
	title      string
	upvotes    int64
	url        string
}

type Comment struct {
	author      string
	createdUtc  int64
	depth       int
	id          string
	jsonReplies []gjson.Result
	message     string
	replies     []*Comment
	score       int64
	thread      *Thread
}
