package fourchan

import (
	"github.com/moshee/go-4chan-api/api"
)

func Fetch(rawUrl string, width, leftPadding uint) (*Thread, error) {
	threadId, board, err := parseUrl(rawUrl)
	if err != nil {
		return nil, err
	}
	apiThread, err := api.GetThread(board, int64(threadId))
	if err != nil {
		return nil, err
	}

	thread := toThread(apiThread)
	thread.width = width
	thread.leftPadding = leftPadding
	thread.op.thread = thread

	return thread, nil
}
