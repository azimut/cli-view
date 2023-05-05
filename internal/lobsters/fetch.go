package lobsters

import (
	"errors"
	"time"

	"github.com/azimut/cli-view/internal/fetch"
	"github.com/tidwall/gjson"
)

var ErrInvalidJson = errors.New("invalid json response")

func Fetch(rawUrl, userAgent string, timeout time.Duration) (*Thread, error) {
	jsonUrl, err := effectiveUrl(rawUrl)
	if err != nil {
		return nil, err
	}
	jsonRaw, err := fetch.Fetch(jsonUrl, userAgent, timeout)
	if err != nil {
		return nil, err
	}
	if !gjson.Valid(jsonRaw) {
		return nil, ErrInvalidJson
	}

	thread := Thread{}
	thread.op = Op{
		createdAt: gjson.Get(jsonRaw, "created_at").Time(),
		message:   gjson.Get(jsonRaw, "description").String(),
		nComments: int(gjson.Get(jsonRaw, "comment_count").Int()),
		score:     int(gjson.Get(jsonRaw, "score").Int()),
		self:      rawUrl,
		url:       gjson.Get(jsonRaw, "url").String(),
		id:        gjson.Get(jsonRaw, "short_id").String(),
		title:     gjson.Get(jsonRaw, "title").String(),
		username:  gjson.Get(jsonRaw, "submitter_user.username").String(),
		thread:    &thread,
	}
	comments := gjson.Get(jsonRaw, "comments").Array()
	for _, comment := range comments {
		thread.insert(Comment{
			createdAt: comment.Get("created_at").Time(),
			id:        comment.Get("short_id").String(),
			message:   comment.Get("comment").String(),
			parent:    comment.Get("parent_comment").String(),
			score:     int(comment.Get("score").Int()),
			thread:    &thread,
			username:  comment.Get("commenting_user.username").String(),
		})
	}
	return &thread, nil
}
