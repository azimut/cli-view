package hackernews

import (
	"sync"
	"time"

	"github.com/caser/gophernews"
	"github.com/pkg/errors"
)

type result struct {
	comment gophernews.Comment
	err     error
}

func Fetch(
	rawUrl string,
	timeout time.Duration,
	maxComments uint,
	nWorkers uint,
) (*Thread, error) {

	url, storyId, err := effectiveUrl(rawUrl)
	if err != nil {
		return nil, err
	}

	client := gophernews.NewClient()
	story, err := fetchStory(client, storyId)
	if err != nil {
		return nil, err
	}

	thread := Thread{}
	op := Op{
		date:      unix2time(story.Time()),
		kids:      story.Kids(),
		ncomments: len(story.Kids()), // TODO: this only gets the direct replies
		score:     story.Score(),
		selfUrl:   url,
		text:      story.Text(),
		thread:    &thread,
		title:     story.Title(),
		url:       story.URL(),
		user:      story.By(),
	}

	commentIds := op.kids
	commentIds = commentIds[:min(len(commentIds), int(maxComments))]
	comments := fetchComments(commentIds, nWorkers, &thread) // TODO: error
	thread.op = op
	thread.comments = comments

	return &thread, nil
}

func fetchStory(client *gophernews.Client, id int) (gophernews.Item, error) {
	item, err := client.GetItem(id)
	if err != nil {
		return nil, err
	}
	if item.Type() != "story" {
		return nil, errors.New("invalid type returned while story expected")
	}
	return item, nil
}

func fetchComments(commentIds []int, nWorkers uint, thread *Thread) []Comment {
	if len(commentIds) == 0 {
		return nil
	}
	idsCh := make(chan int, 5)
	commentsCh := make(chan result, 5)
	var wg sync.WaitGroup
	wg.Add(len(commentIds))
	go sendWork(commentIds, idsCh)
	go closeAfterWait(&wg, commentsCh)
	for i := 0; i < int(nWorkers); i++ {
		go worker(&wg, idsCh, commentsCh)
	}
	return collector(&wg, idsCh, commentsCh, thread)
}

func closeAfterWait(wg *sync.WaitGroup, commentsCh chan<- result) {
	wg.Wait()
	close(commentsCh)
}

func sendWork(ids []int, input chan<- int) {
	for _, id := range ids {
		input <- id
	}
	close(input)
}

func worker(wg *sync.WaitGroup, commentsChan <-chan int, output chan<- result) {
	client := gophernews.NewClient()
	for commentId := range commentsChan {
		comment, err := client.GetComment(commentId)
		if err != nil {
			output <- result{err: err}
		}
		output <- result{comment: comment}
		wg.Done()
	}
}

func isDeleted(comment Comment) bool {
	return comment.msg == "" || comment.msg == "[flagged]" || comment.msg == "[dead]"
}

func isChildless(comment Comment) bool {
	return len(comment.kids) == 0
}

func collector(
	wg *sync.WaitGroup,
	commentsCh chan<- int,
	responseCh <-chan result,
	thread *Thread,
) []Comment {
	var comments []Comment

	for response := range responseCh {

		if response.err != nil {
			continue
		}

		comment := newComment(response.comment, thread)
		if isDeleted(comment) {
			continue
		}
		comments = append(comments, comment)

		// for _, id := range comment.Kids {
		// 	commentsCh <- id
		// }
	}
	return comments
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
