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

	op := toOp(parsedJson)
	op.thread = &thread
	thread.op = op

	rawComments := parsedJson.Get("post_stream.posts").Array()[1:]
	for _, rawComment := range rawComments {
		if isAnAction(rawComment) {
			continue
		}
		comment := toComment(rawComment)
		comment.thread = &thread
		thread.insertComment(comment)
	}

	return &thread, nil
}

func isAnAction(result gjson.Result) bool {
	return result.Get("cooked").String() == "" && result.Get("action_code").String() != ""
}

func toOp(resultJson gjson.Result) Op {
	resultOp := resultJson.Get("post_stream.posts.0")
	return Op{
		id:        int(resultOp.Get("post_number").Int()),
		author:    resultOp.Get("username").String(),
		message:   resultOp.Get("cooked").String(),
		createdAt: resultJson.Get("created_at").Time(),
		title:     resultJson.Get("title").String(),
	}
}

func toComment(resultComment gjson.Result) Comment {
	// TODO: depth, replies
	return Comment{
		author:    resultComment.Get("username").String(),
		createdAt: resultComment.Get("created_at").Time(),
		id:        int(resultComment.Get("post_number").Int()),
		message:   resultComment.Get("cooked").String(),
		parentId:  int(resultComment.Get("reply_to_post_number").Int()),
	}
}
