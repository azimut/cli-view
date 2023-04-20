package fourchan

import (
	"testing"
)

func TestFindPostInThread(t *testing.T) {
	testTread := Thread{
		op: Op{
			id: 1,
		},
		posts: []Post{
			{id: 2},
			{id: 3,
				replies: []Post{
					{id: 4,
						replies: []Post{
							{id: 5},
							{id: 6,
								replies: []Post{
									{id: 7},
									{id: 8},
								},
							},
						}},
				}},
		},
	}
	needle := 8
	foundPost, depth, foundit := testTread.find(needle)
	if foundit != true {
		t.Errorf("could not found id %d on thread", needle)
	}

	expected := 4
	got := depth
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}

	expected = needle
	got = foundPost.id
	if expected != got {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func TestFindPost(t *testing.T) {
	testPost := Post{id: 4,
		replies: []Post{
			{id: 5},
			{id: 6,
				replies: []Post{
					{id: 7},
					{id: 8},
				},
			},
		}}

	_, depth := testPost.find(8, 0)

	expected := 2
	got := depth
	if got != expected {
		t.Errorf("got %d expected %d", got, expected)
	}
}
