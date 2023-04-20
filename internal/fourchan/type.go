package fourchan

import (
	"time"
)

type Thread struct {
	closed bool
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
	id         int
	parentId   int
	replies    []Post
	subject    string // NOTE: mainly used for debug on testing
}

type Attachment struct {
	filename string
	id       string
	url      string
}

// insert we assume the parentId was properly set outside
func (thread *Thread) insert(post Post) {
	// we stop here if is a direct response
	if post.parentId == 0 || post.parentId == thread.op.id {
		thread.posts = append(thread.posts, post)
		return
	}
	// try to find parentId on thread
	parentPost, depth, found := thread.find(post.parentId)
	if found {
		post.depth = depth
		newReplies := append(parentPost.replies, post)
		parentPost.replies = newReplies
	} else {
		thread.posts = append(thread.posts, post) // TODO: fallback
	}
}

func (thread *Thread) find(needlePostId int) (*Post, int, bool) {
	for i := range thread.posts {
		foundPost, depth := thread.posts[i].find(needlePostId, 1)
		if foundPost != nil {
			return foundPost, depth, true
		}
	}
	return nil, 0, false
}

func (post *Post) find(needlePostId, depth int) (*Post, int) {
	if post.id == needlePostId {
		return post, depth
	}
	for i := range post.replies {
		foundPost, newDepth := post.replies[i].find(needlePostId, depth+1)
		if foundPost != nil {
			return foundPost, newDepth
		}
	}
	return nil, 0
}
