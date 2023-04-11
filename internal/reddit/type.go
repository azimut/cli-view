package reddit

import "github.com/tidwall/gjson"

type Thread struct {
	op       Op
	comments []Comment
}

type Op struct {
	author     string
	createdUTC int64
	nComments  int64
	self       string
	selftext   string
	title      string
	upvotes    int64
	url        string
}

type Comment struct {
	author      string
	createdUtc  int64
	depth       int64
	isOp        bool
	message     string
	replies     []Comment
	jsonReplies []gjson.Result
	score       int64
}
