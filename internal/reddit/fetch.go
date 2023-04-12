package reddit

import (
	"time"

	"github.com/azimut/cli-view/internal/fetch"
	"github.com/tidwall/gjson"
)

func Fetch(rawUrl string, ua string, timeout time.Duration) (*Thread, error) {
	url, err := effectiveUrl(rawUrl)
	if err != nil {
		return nil, err
	}

	rawJson, err := fetch.Fetch(url, ua, timeout)
	if err != nil {
		return nil, err
	}

	if !gjson.Valid(rawJson) {
		return nil, err
	}

	thread := toThread(rawJson, rawUrl)
	return thread, nil
}

func toThread(rawJson string, rawUrl string) *Thread {
	post := gjson.Get(rawJson, "0.data.children.0.data")
	op := Op{
		author:     post.Get("author").String(),
		createdUTC: post.Get("created_utc").Int(),
		nComments:  post.Get("num_comments").Int(),
		self:       rawUrl,
		selftext:   post.Get("selftext").String(),
		title:      post.Get("title").String(),
		upvotes:    post.Get("ups").Int(),
		url:        post.Get("url").String(),
	}

	if op.nComments <= 0 {
		return &Thread{op: op}
	}

	var comments []Comment
	var json_comments []gjson.Result

	json_comments = gjson.Get(rawJson, "1.data.children.#.data").Array()
	for _, json_comment := range json_comments {
		comment := toComment(json_comment, op.author)
		// TODO: is probably a "More..." link
		if comment.author == "" {
			continue
		}
		addReplies(&comment, op.author)
		comments = append(comments, comment)
	}
	return &Thread{
		op:       op,
		comments: comments,
	}
}

func addReplies(parentComment *Comment, opAuthorName string) {
	for _, jsonReply := range parentComment.jsonReplies {
		comment := toComment(jsonReply, opAuthorName)
		// TODO: is probably a "More..." link
		if comment.author == "" {
			continue
		}
		parentComment.replies = append(parentComment.replies, comment)
		addReplies(&comment, opAuthorName)
	}
}

func toComment(jsonComment gjson.Result, opAuthorName string) Comment {
	author := jsonComment.Get("author").String()
	return Comment{
		author:      author,
		createdUtc:  jsonComment.Get("created_utc").Int(),
		depth:       jsonComment.Get("depth").Int(),
		isOp:        opAuthorName == author,
		jsonReplies: jsonComment.Get("replies.data.children.#.data").Array(),
		message:     jsonComment.Get("body").String(),
		score:       jsonComment.Get("score").Int(),
	}
}