package hackernews

import (
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func MakeComments(doc *goquery.Document) []*Comment {
	prev := &Comment{}
	var parents []*Comment
	var comments []*Comment
	doc.Find("table.comment-tree tr.comtr").Each(func(i int, sel *goquery.Selection) {
		current := NewComment(sel)
		if current.isChildOf(prev) {
			prev.Childs = append(prev.Childs, current)
			parents = append(parents, prev)
		}
		if current.isSiblingOf(prev) {
			if len(parents) > 0 {
				parents[len(parents)-1].Childs = append(parents[len(parents)-1].Childs, current)
			} else {
				comments = append(comments, current)
			}
		}
		if current.isAncestorOf(prev) {
			diff := prev.indent - current.indent
			parents = parents[:len(parents)-diff]
			if len(parents) > 0 {
				parents[len(parents)-1].Childs = append(parents[len(parents)-1].Childs, current)
			} else {
				comments = append(comments, current)
			}
		}
		prev = current
	})
	return comments
}

func NewComment(sel *goquery.Selection) *Comment {
	return &Comment{
		id:  commentId(sel),
		msg: commentMsg(sel),
		//togg: commentTogg(sel),
		user: commentUser(sel),
		//date:     commentDate(sel),
		indent: commentIndent(sel),
	}
}

func (current *Comment) isChildOf(other *Comment) bool {
	if current.indent > other.indent {
		return true
	}
	return false
}

func (current *Comment) isSiblingOf(other *Comment) bool {
	if current.indent == other.indent {
		return true
	}
	return false
}

func (current *Comment) isAncestorOf(other *Comment) bool {
	if current.indent < other.indent {
		return true
	}
	return false
}

func commentTogg(sel *goquery.Selection) int {
	rawTogg, exists := sel.Find("a.togg").Attr("n")
	if !exists {
		panic("no toggle n on comment")
	}
	togg, err := strconv.Atoi(rawTogg)
	if err != nil {
		panic(err)
	}
	return togg
}

func commentIndent(sel *goquery.Selection) int {
	rawindent, exists := sel.Find("td.ind").Attr("indent")
	if !exists {
		panic("no indent for comment")
	}
	indent, err := strconv.Atoi(rawindent)
	if err != nil {
		panic(err)
	}
	return indent
}

func commentMsg(sel *goquery.Selection) string {
	return sel.Find("span.commtext").Text()
}

func commentId(sel *goquery.Selection) int {
	rawid, exists := sel.Attr("id")
	if !exists {
		panic("comment id not found")
	}
	id, err := strconv.Atoi(rawid)
	if err != nil {
		panic(err)
	}
	return id
}

func commentUser(sel *goquery.Selection) string {
	return sel.Find("a.hnuser").Text()
}

func commentDate(sel *goquery.Selection) time.Time {
	rawdate, exists := sel.Find("span.age").Attr("title")
	if !exists {
		panic("could not find span.age title")
	}
	date, err := time.Parse("%-%M-%DT%h:%m:%s", rawdate)
	if err != nil {
		panic(err)
	}
	return date
}
