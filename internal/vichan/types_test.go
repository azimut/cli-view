package vichan

import "testing"

func TestExplode(t *testing.T) {
	testComments := []Comment{
		{
			subject: "1 plaintext comment",
			message: "Its crazy you can easily do stuff like this with libraries like mesa on python",
		},
		{
			subject: "1 comment with a random link",
			message: `I started building a systematic gallery of cellular automata: <a href="https://cenysor.neocities.org/programming/cells/gallery/cells_gallery" rel="nofollow" target="_blank">https://cenysor.neocities.org/programming/cells/gallery/cells_gallery</a>`,
		},
		{
			subject: "1 comment with 1 reply on the same thread",
			message: `<a onclick="highlightReply('35080', event);" href="/λ/res/35065.html#35080">&gt;&gt;35080</a><br/>A lot of these seem to produce stripes or checkerboard patterns.`,
		},
		{
			subject: "1 comment with <span> quote",
			message: `<span class="quote">&gt;32 - Chan aggregator</span><br/>Pretty interested at trying my hands at this. I bounce between alot of chans myself. Anyone have material I can use to study this topic?`,
		},
		{
			subject: "2 comments, implicit reply to OP, 1 reply to comment in thread",
			message: `Thought a little bit about something I posted earlier: CL already has type parameters (e.g. (deftype fixnum-car-cons () '(cons fixnum *))) but no way to parametrize classes. And I really want it because I'm making a hash map using chained hashing, so the bucket type (list, vector or BST) should be customizable by the user for its performance characteristics.<br/>Instead, I'll have to use some godless macro abominations and define different classes for hash-map/list, hash-map/vector, etc...<br/><a onclick="highlightReply('31416', event);" href="/λ/res/31415.html#31416">&gt;&gt;31416</a><br/>I'd say we understand each other now. You want spinlocks to avoid the kernel machinery (sleep/wake up, mainly) cost, but with an io_uring like mechanism of shared memory to tell the scheduler to not bother you. Interesting.`,
		},
		{
			subject: "1 comment, 2 replies with the same comment",
			message: `<a onclick="highlightReply('31635', event);" href="/λ/res/31415.html#31635">&gt;&gt;31635</a><br/><a onclick="highlightReply('31628', event);" href="/λ/res/31415.html#31628">&gt;&gt;31628</a><br/>Ok, lainanons. I appreciate your help and support. But maybe lisp is not just for me. Maybe I'm too dumb to use lisps. Maybe I'm too picky and I shall stay at learning mainstream soykaf like python. Python is kinda ok for most of the stuff I regularly do. But it's boring. I just wanted to find my favorite language. I've been looking for something special. And nothing got me like lisp. Nor C-thing, nor haskell, nor java(script). You know lisp is kinda romantic. Lisp hacker sounds cool. And I really appreciate you code lisp despite its unpopularity. But what I like more is that  you've found a language you love, in which you really enjoy programming.`,
		},
		{
			subject: "1 comment, 3 replies with the same comment",
			message: `<a onclick="highlightReply('31699', event);" href="/λ/res/31415.html#31699">&gt;&gt;31699</a><br/><a onclick="highlightReply('31635', event);" href="/λ/res/31415.html#31635">&gt;&gt;31635</a><br/><a onclick="highlightReply('31628', event);" href="/λ/res/31415.html#31628">&gt;&gt;31628</a><br/>Ok, lainanons. I appreciate your help and support. But maybe lisp is not just for me. Maybe I'm too dumb to use lisps. Maybe I'm too picky and I shall stay at learning mainstream soykaf like python. Python is kinda ok for most of the stuff I regularly do. But it's boring. I just wanted to find my favorite language. I've been looking for something special. And nothing got me like lisp. Nor C-thing, nor haskell, nor java(script). You know lisp is kinda romantic. Lisp hacker sounds cool. And I really appreciate you code lisp despite its unpopularity. But what I like more is that  you've found a language you love, in which you really enjoy programming.`,
		},
		{
			subject: "2 messages, with 2 different replies",
			message: `<a onclick="highlightReply('31646', event);" href="/λ/res/31415.html#31646">&gt;&gt;31646</a><br/><span class="quote">&gt;but still not at the CMUCL/SBCL level</span><br/>Except Java running on JVM is the same as SBCL in terms of speed.<br/>I'm not saying JIT and AOT are exactly the same, but they are very similar in a general sense.<br/>SBCL is more like a JIT than AOT. It's just the way SBCL uses images means you can effectively save the JITed code as if it where AOT, but the way it works under the hood is inevitably going to be far closer to a JIT compiler.<br/><br/><a onclick="highlightReply('31649', event);" href="/λ/res/31415.html#31649">&gt;&gt;31649</a><br/><span class="quote">&gt;so what is it? maybe I still haven't got that</span><br/>I don't know what you don't like. My point is that the things on your list of complaints are all very small compared to the benefit of "the most powerful language ever existed".<br/>There must be something else that you do not like about it to justify this overall negative perception.<br/>Or alternatively, one could conclude that "powerful" isn't really a very useful trait for a programming language. A lot of programming is inherently difficult and languages can only do so much to help, thus you are disappointed when all this power isn't actually making your life that much easier.`,
		},
	}
	expecteds := []int{1, 1, 1, 1, 2, 1, 1, 2} // 2
	for i, comment := range testComments {
		comments, err := comment.explode()
		if err != nil {
			t.Errorf("%v", err)
		}
		got := len(comments)
		expected := expecteds[i]
		if expected != got {
			t.Errorf("got %d expected %d - %s", got, expected, comment.subject)
		}
	}
}

func TestAllEmptyButLast(t *testing.T) {
	testReplies := []struct {
		replies  []string
		expected bool
	}{
		{[]string{"", "", "foo"}, true},
		{[]string{"foo", "bar"}, false},
		{[]string{"", "foo"}, true},
		{[]string{}, false},
	}
	for _, reply := range testReplies {
		got := allEmptyButLast(reply.replies)
		expected := reply.expected
		if expected != got {
			t.Errorf(
				"got %t expected %t - (%d)%+v",
				got,
				expected,
				len(reply.replies),
				reply.replies,
			)
		}
	}
}

func TestAnyEmptyString(t *testing.T) {
	testSlices := [][]string{
		{"foo", "bar"},
		{"foo", "", "baz"},
		{""},
	}
	testExpected := []bool{false, true, true}
	for i, slice := range testSlices {
		expected := testExpected[i]
		got := anyEmptyString(slice)
		if expected != got {
			t.Errorf("got %t expected %t", got, expected)
		}
	}
}

func TestGetCommentThread(t *testing.T) {
	testThread := Thread{
		comments: []Comment{
			{
				id: 1,
				replies: []Comment{
					{id: 2},
					{id: 3},
				},
			},
		},
	}
	a := testThread.getComment(3)
	b := testThread.getComment(3)
	if a != b {
		t.Errorf("pointers do not match %p != %p", a, b)
	}
}

func TestGetComment(t *testing.T) {
	testComment := Comment{
		id: 1,
		replies: []Comment{
			{id: 2},
			{id: 3},
		},
	}
	a := testComment.getComment(3)
	b := testComment.getComment(3)
	if a != b {
		t.Errorf("pointers do not match %p != %p", a, b)
	}
}
