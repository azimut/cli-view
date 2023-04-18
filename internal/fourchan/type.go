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
	id         int
	parentId   int
	replies    []Post
	subject    string
}

type Attachment struct {
	filename string
	id       string
	url      string
}

func (thread *Thread) insert(post Post) {
	// either parentId is not set, or is the OP
	if post.parentId == 0 || post.parentId == thread.op.id {
		thread.posts = append(thread.posts, post)
		return
	}
	// try to find parentId on thread
	parentPost, found := thread.find(post.parentId)
	if found {
		parentPost.replies = append(parentPost.replies, post)
	} else {
		thread.posts = append(thread.posts, post) // TODO: fallback
	}
}

func (thread *Thread) find(postId int) (*Post, bool) {
	for _, post := range thread.posts {
		foundPost := post.find(postId)
		if foundPost != nil {
			return foundPost, true
		}
	}
	return nil, false
}

func (post *Post) find(postId int) *Post {
	if post.id == postId {
		return post
	}
	for _, reply := range post.replies {
		found := reply.find(postId)
		if found != nil {
			return found
		}
	}
	return nil
}
