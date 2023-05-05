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

func (t *Thread) insert(comment Comment) {
	if comment.parent == "" {
		t.comments = append(t.comments, comment)
	} else {
		parentComment := t.findComment(comment.parent)
		comment.depth = parentComment.depth + 1
		parentComment.replies = append(parentComment.replies, comment)
	}
}

func (t *Thread) findComment(id string) *Comment {
	for i := range t.comments {
		if found := t.comments[i].findComment(id); found != nil {
			return found
		}
	}
	return nil
}

func (c *Comment) findComment(id string) *Comment {
	if c.id == id {
		return c
	}
	for i := range c.replies {
		if found := c.replies[i].findComment(id); found != nil {
			return found
		}
	}
	return nil
}
