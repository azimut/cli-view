package hackernews

import (
	"errors"
	"sync"
	"time"

	"github.com/caser/gophernews"
)

const NUM_OF_WORKERS = 3

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

func Fetch(rawUrl, ua string, timeout time.Duration) (doc *Op, c *[]Comment, err error) {
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
	comments := fetchComments(story.Kids) // TODO: error
	return &op, &comments, nil
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

func fetchComments(comments []int) []Comment {
	idsChan := make(chan int, 5)
	commentsCh := make(chan result, 5)
	var wg sync.WaitGroup
	go sendWork(&wg, comments, idsChan)
	go closeAfterWait(&wg, commentsCh)
	spawnWorkers(&wg, idsChan, commentsCh)
	return collector(&wg, idsChan, commentsCh)
}

func closeAfterWait(wg *sync.WaitGroup, commentsCh chan<- result) {
	wg.Wait()
	close(commentsCh)
}

func sendWork(wg *sync.WaitGroup, ids []int, input chan<- int) {
	for _, id := range ids {
		input <- id
		wg.Add(1)
	}
	close(input)
}

func spawnWorkers(wg *sync.WaitGroup, in <-chan int, out chan<- result) {
	for i := 0; i < NUM_OF_WORKERS; i++ {
		go worker(wg, in, out)
	}
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
		state = append(state, Comment{
			id:   c.ID,
			msg:  c.Text,
			user: c.By,
			kids: c.Kids,
		})
		// for _, id := range c.Kids {
		// 	commentsCh <- id
		// }
	}
	return
}
