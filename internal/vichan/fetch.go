package vichan

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/azimut/cli-view/internal/fetch"
	"github.com/azimut/cli-view/internal/vichan/jsonmodel"
)

func Fetch(
	rawUrl, userAgent string,
	width, leftPadding uint,
	timeout time.Duration,
) (*Thread, error) {
	effectiveUrl, err := parseUrl(rawUrl)
	if err != nil {
		return nil, err
	}

	rawJson, err := fetch.Fetch(effectiveUrl, userAgent, timeout)
	if err != nil {
		return nil, err
	}

	var vichanThread jsonmodel.Thread
	if err = json.Unmarshal([]byte(rawJson), &vichanThread); err != nil {
		return nil, err
	}

	thread, err := toThread(vichanThread)
	if err != nil {
		return nil, err
	}

	thread.op.thread = thread
	thread.leftPadding = leftPadding
	thread.width = width
	thread.url = rawUrl

	return thread, nil
}

func toThread(vichanThread jsonmodel.Thread) (*Thread, error) {
	if len(vichanThread.Posts) == 0 {
		return nil, errors.New("malformed post, doesn't have OP")
	}

	vichanOp := vichanThread.Posts[0]
	op := Op{
		createdAt: time.Unix(int64(vichanOp.Time), 0),
		id:        vichanOp.No,
		message:   vichanOp.Comment,
		title:     vichanOp.Title,
	}

	thread := &Thread{op: op}
	for _, post := range vichanThread.Posts[1:] {
		newComment := toComment(post)
		newComment.thread = thread
		newComments, err := newComment.explode()
		if err != nil {
			return nil, err
		}
		for _, comment := range newComments {
			thread.insert(comment)
		}
	}
	return thread, nil
}

func toComment(message jsonmodel.Message) Comment {
	return Comment{
		author:    message.Author,
		createdAt: time.Unix(int64(message.Time), 0),
		id:        message.No,
		message:   message.Comment,
	}
}
