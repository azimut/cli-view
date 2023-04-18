package fourchan

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/moshee/go-4chan-api/api"
)

var reQuote = regexp.MustCompile(`<a href="[^"]+" class="quotelink">&gt;&gt;[0-9]+</a>[ ]*<br>`)

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
		newPost := Post{
			attachment: getAttachment(apiPost),
			comment:    apiPost.Comment,
			created:    apiPost.Time,
			id:         int(apiPost.Id),
			parentId:   thread.op.id,
			subject:    apiPost.Subject,
		}
		for _, post := range explodePost(newPost) {
			thread.insert(post)
		}
	}
	return &thread
}

// explodePost explodes based on "quotelink"
func explodePost(post Post) (posts []Post) {
	findings := reQuote.FindAllString(post.comment, -1)
	replies := reQuote.Split(post.comment, -1)

	// Whole post is "not replying"
	if len(findings) == 0 {
		return append(posts, post)
	}

	// At least part of the post is "not replying"
	if len(replies) > 0 && replies[0] != "" {
		newPost := post
		newPost.comment = replies[0]
		posts = append(posts, newPost)
	}

	// Remove (by now processed) not reply part
	replies = replies[1:]
	if len(replies) == 0 {
		return
	}

	// Add simple 1/1 reply/comment
	if (!containsEmptyString(replies)) && len(replies) == len(findings) {
		for i, reply := range replies {
			id := getParentId(findings[i])
			if id == 0 {
				continue // TODO: handle external links?
			}
			newPost := post
			newPost.comment = reply
			newPost.parentId = id
			posts = append(posts, newPost)
		}
		return
	}

	fmt.Println("---- Findings ", len(findings))
	for _, finding := range findings {
		fmt.Println(finding)
	}
	fmt.Println("---- Comment <" + post.subject + ">")
	fmt.Println(post.comment)
	fmt.Println("---- Replies ", len(replies))
	for _, reply := range replies {
		fmt.Println("'" + reply + "'")
	}
	fmt.Println("--------------------------------------------------")
	return
}

func getParentId(finding string) int {
	if strings.Contains(finding, `/g/thread/`) {
		return 0
	}
	begin := strings.Index(finding, "#p")
	if begin == -1 {
		return 0
	}
	rawId := finding[begin+2:]
	rawId = rawId[0:strings.Index(rawId, `"`)]
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return 0
	}
	return id
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

func containsEmptyString(xs []string) bool {
	for _, x := range xs {
		if x == "" {
			return true
		}
	}
	return false
}
