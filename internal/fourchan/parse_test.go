package fourchan

import (
	"testing"
	"time"

	"github.com/moshee/go-4chan-api/api"
)

var testOp = &api.Post{
	Comment: "OP plaintext",
	Id:      92834152,
	Subject: "my op subject",
	Time:    time.Now(),
}

var testThread = &api.Thread{
	Board: "g",
	OP:    testOp,
	Posts: []*api.Post{
		testOp,
		{
			Subject: "same thread SINGLE reply (to OP) and a link",
			Comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br>Yes https://jetbra.in/s`,
			Id:      92835905,
			Time:    time.Now(),
			File: &api.File{
				Ext:  ".jpg",
				Id:   1653870571633,
				Name: "wacky",
			},
			Thread: &api.Thread{Board: "g"}, // NOTE: fake value
		},
		{
			Subject: "same thread SINGLE reply and a link",
			Comment: ``,
			Id:      92835905,
			Time:    time.Now(),
			File: &api.File{
				Ext:  ".jpg",
				Id:   1653870571633,
				Name: "ayaya",
			},
			Thread: &api.Thread{Board: "g"}, // NOTE: fake value
		},
		{
			Subject: "same thread SINGLE reply, no comment but the link to parent, no <br>",
			Comment: `<a href="#p92835905" class="quotelink">&gt;&gt;92835905</a>`,
			Id:      92838669,
			Time:    time.Now(),
			File: &api.File{
				Ext:  ".jpg",
				Id:   1653870571633,
				Name: "miyanohype",
			},
			Thread: &api.Thread{Board: "g"}, // NOTE: fake value
		},
		{
			Subject: "same thread SINGLE reply, and non linked reply to OP",
			Comment: `No YOU<br><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>Yes https://jetbra.in/s`,
			Id:      92835911,
			Time:    time.Now(),
		},
		{
			Subject: "plaintext comment",
			Comment: "hey whats up",
			Id:      92838633,
			Time:    time.Now(),
		},
		{
			Subject: "same thread 2 replies",
			Comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br>paying for free software?<br><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>based`,
			Id:      92838617,
			Time:    time.Now(),
		},
		{
			Subject: "same thread 2 replies, with the same message ",
			Comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>based`,
			Id:      92838699,
			Time:    time.Now(),
		},
		{
			Subject: "same thread 3 replies",
			Comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br>paying for free software?<br><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>based<br><br><a href="#p92838617" class="quotelink">&gt;&gt;92838617</a><br>paying for free software?<br><br>`,
			Id:      92838617,
			Time:    time.Now(),
		},
	},
}

func TestOp(t *testing.T) {
	thread := toThread(testThread)
	got := thread.op.id
	expected := 92834152
	if got != expected {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func TestExplodeNPosts(t *testing.T) {
	var (
		expected int
		got      int
	)
	testPosts := []Post{
		{
			subject: "comment, 1 response",
			comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br>Yes https://jetbra.in/s`,
		},
		{
			subject: "comment, no reply",
			comment: "hey whats up",
		},
		{
			subject: "empty comment, no reply",
		},
		{
			subject: "empty comment, 1 reply",
			comment: `<a href="#p92835905" class="quotelink">&gt;&gt;92835905</a>`,
		},
		{
			subject: "2 replies",
			comment: `No YOU<br><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>Yes https://jetbra.in/s`,
		},
		{
			subject: "3 replies",
			comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br>paying for free software?<br><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>based<br><br><a href="#p92838617" class="quotelink">&gt;&gt;92838617</a><br>paying for free software?<br><br>`,
		},
		{
			subject: "2 replies, with the same message",
			comment: `<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br><a href="#p92835905" class="quotelink">&gt;&gt;92835905</a><br>based`,
		},
	}
	testNrs := []int{1, 1, 1, 1, 2, 3, 1}
	for i, rawPost := range testPosts {
		posts := explodePost(rawPost)
		expected = testNrs[i]
		got = len(posts)
		if expected != got {
			t.Errorf("got %d expected %d - on %s - %+v", got, expected, rawPost.subject, posts)
		}
	}
}

func TestNParents(t *testing.T) {
	thread := toThread(testThread)
	got := len(thread.posts)
	expected := 8
	if got != expected {
		t.Errorf("got %d expected %d", got, expected)
	}
}

func TestGetParentId(t *testing.T) {
	findings := map[string]int{
		`<a href="/g/thread/92865685#p92866880" class="quotelink">&gt;&gt;92866880</a> <br>`: 0,
		`<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a><br>`:                    92834152,
		`<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a> <br>`:                   92834152,
		`<a href="#p92834152" class="quotelink">&gt;&gt;92834152</a>`:                        92834152,
		``: 0,
	}
	for finding, expected := range findings {
		got := getParentId(finding)
		if expected != got {
			t.Errorf("got %d expected %d", got, expected)
		}
	}
}
