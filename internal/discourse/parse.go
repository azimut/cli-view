package discourse

import (
	"errors"

	"github.com/tidwall/gjson"
)

func toThread(rawJson, rawUrl string) (*Thread, error) {
	if !gjson.Valid(rawJson) {
		return nil, errors.New("invalid json response")
	}
	parsedJson := gjson.Parse(rawJson)

	var thread Thread
	thread.id = int(parsedJson.Get("id").Int())
	thread.url = rawUrl

	opPost := parsedJson.Get("post_stream.posts.0")
	thread.op = Op{
		id:        int(opPost.Get("id").Int()),
		author:    opPost.Get("username").String(),
		message:   opPost.Get("cooked").String(),
		createdAt: parsedJson.Get("created_at").Time(),
		title:     parsedJson.Get("title").String(),
		thread:    &thread,
	}
	thread.comments = toComments(parsedJson.Get("post_stream.posts").Array()[1:], &thread)
	return &thread, nil
}

func toComments(results []gjson.Result, thread *Thread) (comments []Comment) {
	for _, result := range results {
		if isAnAction(result) {
			continue
		}
		comments = append(comments, toComment(result, thread))
	}
	return
}

func isAnAction(result gjson.Result) bool {
	return result.Get("cooked").String() == "" && result.Get("action_code").String() != ""
}

func toComment(result gjson.Result, thread *Thread) Comment {
	// TODO: depth, replies
	return Comment{
		author:    result.Get("username").String(),
		createdAt: result.Get("created_at").Time(),
		id:        int(result.Get("id").Int()),
		message:   result.Get("cooked").String(),
		thread:    thread,
	}
}
