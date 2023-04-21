package hackernews

import (
	"errors"
	"sync"
	"time"

	"github.com/caser/gophernews"
)

func Fetch(
	rawUrl string,
	timeout time.Duration,
	maxComments int,
	nWorkers uint,
) (doc Op, comments []Comment, err error) {

	url, storyId, err := effectiveUrl(rawUrl)
	if err != nil {
		return Op{}, nil, err
	}

	client := gophernews.NewClient()
	story, err := fetchStory(client, storyId)
	if err != nil {
		return Op{}, nil, err
	}

	ids := story.Kids
	maxComments = min(len(ids), maxComments)
	if maxComments > 0 {
		ids = ids[:maxComments]
		comments = fetchComments(ids, nWorkers) // TODO: error
	}

	op := newOp(story, url)
	return op, comments, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fetchStory(client *gophernews.Client, id int) (*gophernews.Story, error) {
	story, err := client.GetStory(id)
	if err != nil {
		return nil, err
	}
	if story.Type != "story" {
		return nil, errors.New("is not a story")
	}
	return &story, nil
}

func fetchComments(comments []int, workers uint) []Comment {
	idsChan := make(chan int, 5)
	commentsCh := make(chan result, 5)
	var wg sync.WaitGroup
	wg.Add(len(comments))
	go sendWork(comments, idsChan)
	go closeAfterWait(&wg, commentsCh)
	for i := 0; i < int(workers); i++ {
		go worker(&wg, idsChan, commentsCh)
	}
	return collector(&wg, idsChan, commentsCh)
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

type result struct {
	comment gophernews.Comment
	err     error
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

func isDeleted(comment gophernews.Comment) bool {
	return comment.Text == "" || comment.Text == "[flagged]" || comment.Text == "[dead]"
}

func isChildless(comment gophernews.Comment) bool {
	return len(comment.Kids) == 0
}

func collector(
	wg *sync.WaitGroup,
	commentsCh chan<- int,
	responseCh <-chan result,
) []Comment {
	var comments []Comment

	for response := range responseCh {

		if response.err != nil {
			continue
		}

		comment := response.comment

		if isDeleted(comment) {
			continue
		}

		comments = append(comments, Comment{
			id:   comment.ID,
			msg:  comment.Text,
			user: comment.By,
			kids: comment.Kids,
			date: unix2time(comment.Time),
		})

		// for _, id := range comment.Kids {
		// 	commentsCh <- id
		// }
	}
	return comments
}
