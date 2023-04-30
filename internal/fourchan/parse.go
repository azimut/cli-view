package fourchan

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jaytaylor/html2text"
	"github.com/moshee/go-4chan-api/api"
)

var reQuote = regexp.MustCompile(`<a href="[^"]+" class="quotelink">&gt;&gt;[0-9]+</a>([ ]*<br>)?`)

func toThread(apiThread *api.Thread) *Thread {
	thread := Thread{
		op: Op{
			attachment: getAttachment(apiThread.OP),
			board:      apiThread.Board,
			comment:    plaintextComment(cleanComment(apiThread.OP.Comment)),
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
			thread:     &thread,
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
	replies := fixComments(
		reQuote.Split(post.comment, -1),
	) // Clean here for later empty string checks

	// Whole post is "not replying"
	if len(findings) == 0 && len(replies) == 1 {
		post.comment = replies[0]
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

	// squash many/1 reply/comment, when no message
	if len(replies) == len(findings) && allEmptyStrings(replies) {
		replies = replies[len(replies)-1:]
		findings = findings[len(findings)-1:]
	}

	// Add simple 1/1 reply/comment messages
	if len(replies) == len(findings) &&
		(len(replies) == 1 || allEmptyButLast(replies) || !containsEmptyString(replies)) {

		// squash many/1 reply/comment into 1/1
		if allEmptyButLast(replies) {
			replies = replies[len(replies)-1:] // keep the last reply
		}

		for i, reply := range replies {
			parentId := getParentId(findings[i])
			if parentId == 0 {
				continue // TODO: handle external links?
			}
			newPost := post
			newPost.comment = reply
			newPost.parentId = parentId
			posts = append(posts, newPost)
		}
		return
	}

	// All that remains are powerusers replying to many with a same message
	fmt.Println("---- Findings ", len(findings))
	for _, finding := range findings {
		fmt.Println(finding)
	}
	fmt.Printf("---- Comment [%d] <%s>\n", post.id, post.subject+">")
	fmt.Println(strings.ReplaceAll(post.comment, "<br>", "<br>\n"))
	fmt.Println("---- Replies ", len(replies))
	for _, reply := range replies {
		fmt.Println("'" + reply + "'")
	}
	fmt.Println("--------------------------------------------------")
	return
}

func allEmptyButLast(replies []string) bool {
	if len(replies) == 0 {
		return false
	}
	for _, reply := range replies[:len(replies)-1] {
		if reply != "" {
			return false
		}
	}
	return replies[len(replies)-1] != ""
}

func fixComments(comments []string) (ret []string) {
	for _, comment := range comments {
		ret = append(ret, plaintextComment(cleanComment(comment)))
	}
	return
}

func plaintextComment(comment string) string {
	comment, err := html2text.FromString(comment, html2text.Options{})
	if err != nil {
		panic(err)
	}
	return comment
}

func cleanComment(comment string) string {
	comment = strings.ReplaceAll(comment, "<wbr>", "") // NOTE: breaks links
	comment = strings.TrimSpace(comment)
	for strings.HasPrefix(comment, "<br>") {
		comment = strings.TrimSpace(strings.TrimPrefix(comment, "<br>"))
	}
	for strings.HasSuffix(comment, "<br>") {
		comment = strings.TrimSpace(strings.TrimSuffix(comment, "<br>"))
	}
	return comment
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

func allEmptyStrings(xs []string) bool {
	for _, x := range xs {
		if x != "" {
			return false
		}
	}
	return true
}
