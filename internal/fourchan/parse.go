package fourchan

import "github.com/moshee/go-4chan-api/api"

func toThread(apiThread *api.Thread) *Thread {
	thread := Thread{
		op: Op{
			attachment: getAttachment(apiThread.OP),
			board:      apiThread.Board,
			comment:    apiThread.OP.Comment,
			created:    apiThread.OP.Time,
			id:         int(apiThread.OP.Id),
			subject:    apiThread.OP.Subject,
		},
	}
	// NOTE: go-4chan-api adds op as the first (aka [0]) post
	if len(apiThread.Posts) <= 1 {
		return &thread
	}
	// TODO: depth + replies tree
	for _, apiPost := range apiThread.Posts[1:] {
		post := Post{
			attachment: getAttachment(apiPost),
			comment:    apiPost.Comment,
			created:    apiPost.Time,
		}
		thread.posts = append(thread.posts, post)
	}
	return &thread
}

func getAttachment(post *api.Post) Attachment {
	if post.ImageURL() == "" {
		return Attachment{}
	}
	return Attachment{
		url:      post.ImageURL(),
		filename: post.File.Name + post.File.Ext,
	}
}
