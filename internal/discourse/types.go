package discourse

import (
	"time"
)

type Thread struct {
	comments     []Comment
	op           Op
	id           int
	url          string
	LineWidth    int
	CommentWidth int
	LeftPadding  int
	ShowDate     bool
	ShowAuthor   bool
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
	parentId  int
	replies   []Comment
	thread    *Thread
}

func (t *Thread) insertComment(comment Comment) {
	if comment.parentId == 1 || comment.parentId == 0 {
		t.comments = append(t.comments, comment)
	} else {
		parentComment := t.findComment(comment.parentId)
		comment.depth = parentComment.depth + 1
		parentComment.replies = append(parentComment.replies, comment)
	}
}

func (t *Thread) findComment(id int) *Comment {
	for i := range t.comments {
		if found := t.comments[i].findComment(id); found != nil {
			return found
		}
	}
	return nil
}

func (c *Comment) findComment(id int) *Comment {
	if id == c.id {
		return c
	}
	for i := range c.replies {
		if found := c.replies[i].findComment(id); found != nil {
			return found
		}
	}
	return nil
}
