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
	// TODO: depth + replies tree
	for _, post := range apiThread.Posts {
		reply := Post{
			attachment: getAttachment(post),
			comment:    post.Comment,
			created:    post.Time,
		}
		thread.posts = append(thread.posts, reply)
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
