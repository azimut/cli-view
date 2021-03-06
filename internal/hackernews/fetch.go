package hackernews

import (
	"errors"
	"sync"
	"time"

	"github.com/caser/gophernews"
)

func unix2time(t int) time.Time {
	return time.Unix(int64(t), 0)
}

func newOp(story *gophernews.Story, selfUrl string) Op {
	return Op{
		url:       story.URL,
		title:     story.Title,
		score:     story.Score,
		user:      story.By,
		date:      unix2time(story.Time),
		ncomments: len(story.Kids), // ?
		selfUrl:   selfUrl,
	}
}

func Fetch(
	rawUrl string,
	timeout time.Duration,
	limit int,
	workers uint,
) (doc *Op, c *[]Comment, err error) {
	url, storyId, err := effectiveUrl(rawUrl)
	if err != nil {
		return nil, nil, err
	}
	client := gophernews.NewClient()
	story, err := fetchStory(client, storyId)
	if err != nil {
		return nil, nil, err
	}
	op := newOp(story, url)
	ids := story.Kids
	limit = min(len(ids), limit)
	if limit > 0 {
		ids = ids[:limit]
	}
	comments := fetchComments(ids, workers) // TODO: error
	return &op, &comments, nil
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

func collector(
	wg *sync.WaitGroup,
	commentsCh chan<- int,
	responseCh <-chan result,

) (state []Comment) {
	for response := range responseCh {
		if response.err != nil {
			continue
		}

		c := response.comment

		if c.Text == "" && len(c.Kids) == 0 {
			continue
		}

		state = append(state, Comment{
			id:   c.ID,
			msg:  c.Text,
			user: c.By,
			kids: c.Kids,
			date: unix2time(c.Time),
		})
		// for _, id := range c.Kids {
		// 	commentsCh <- id
		// }
	}
	return
}
