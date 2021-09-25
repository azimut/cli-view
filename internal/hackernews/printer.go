package hackernews

import (
	"fmt"
	"math"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dustin/go-humanize"
	"jaytaylor.com/html2text"
)

func PrintDoc(doc *goquery.Document) {
	fmt.Println(NewHeader(doc))
	printChilds(NewComments(doc))
}

func printChilds(c []*Comment) {
	for _, value := range c {
		fmt.Println(value)
		printChilds(value.Childs)
	}
}

func (o *Op) String() (ret string) {
	ret += "title: " + o.title + "\n"
	ret += "url: " + o.url + "\n"
	ret += fmt.Sprintf("%s(%d) - %s\n", o.user, o.score, humanize.Time(o.date))
	return
}

func (c *Comment) String() (ret string) {
	indent := c.indent * 5
	msg, err := html2text.FromString(c.msg)
	if err != nil {
		panic(err)
	}
	wrapped := leftPad(msg, 120, indent+1)
	ret += "\n" + wrapped + "\n"
	arrow := ">>> "
	if c.indent > 0 {
		arrow = ">> "
	}
	ret += strings.Repeat(" ", indent) + arrow + c.user + " - " + humanize.Time(c.date)
	return
}

func leftPad(text string, maxWidth int, padding int) string {
	lines := splitByWidthMake(text, maxWidth-padding)
	pad := strings.Repeat(" ", padding)
	for i, line := range lines {
		lines[i] = pad + line
	}
	return strings.Join(lines, "\n")
}

func splitByWidthMake(str string, size int) []string {
	strLength := len(str)
	splitedLength := int(math.Ceil(float64(strLength) / float64(size)))
	splited := make([]string, splitedLength)
	var start, stop int
	for i := 0; i < splitedLength; i += 1 {
		start = i * size
		stop = start + size
		if stop > strLength {
			stop = strLength
		}
		splited[i] = str[start:stop]
	}
	return splited
}
