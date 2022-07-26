package hackernews

import (
	"errors"
	"github.com/caser/gophernews"
	"time"
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

func Fetch(rawUrl, ua string, timeout time.Duration) (doc *Op, err error) {
	url, storyId, err := effectiveUrl(rawUrl)
	if err != nil {
		return nil, err
	}
	client := gophernews.NewClient()
	story, err := fetchStory(client, storyId)
	if err != nil {
		return nil, err
	}
	op := newOp(story, url)
	// comments := fetchComments(story.Kids) // TODO: error
	return &op, nil
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

func fetchComments(comments []int) (state map[int]Comment) {
	storiesCh := make(chan int, 5)
	commentsCh := make(chan result, 5)
	go sendWork(comments, storiesCh)
	spawnWorkers(storiesCh, commentsCh)
	collector(storiesCh, commentsCh, state)
	return state
}

func sendWork(ids []int, input chan<- int) {
	for _, id := range ids {
		input <- id
	}
}

func spawnWorkers(in <-chan int, out chan<- result) {
	for i := 0; i < NUM_OF_WORKERS; i++ {
		go worker(in, out)
	}
}

type result struct {
	comment gophernews.Comment
	err     error
}

func worker(stories <-chan int, output chan<- result) {
	client := gophernews.NewClient()
	for storyId := range stories {
		comment, err := client.GetComment(storyId)
		if err != nil {
			output <- result{err: err}
		}
		output <- result{comment: comment}
	}
}

func collector(commentsCh chan<- int, responseCh <-chan result, state map[int]Comment) {
	for response := range responseCh {
		if response.err != nil {
			continue
		}
		c := response.comment
		state[c.ID] = Comment{
			id:   c.ID,
			msg:  c.Text,
			user: c.By,
			kids: c.Kids,
		}
		for _, id := range c.Kids {
			commentsCh <- id
		}
	}
	close(commentsCh)
}
