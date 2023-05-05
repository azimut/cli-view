package hackernews

import (
	"fmt"
	"html"
	"net/url"
	"regexp"
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/azimut/cli-view/internal/format"
	"github.com/charmbracelet/lipgloss"
	"github.com/dustin/go-humanize"
)

func (thread Thread) String() (ret string) {
	ret += fmt.Sprintln(thread.op)
	for _, comment := range thread.comments {
		ret += fmt.Sprintln(comment)
	}
	return
}

func (o Op) String() (ret string) {
	ret += "\ntitle: " + o.title + "\n"
	if o.url != "" {
		ret += "  url: " + o.url + "\n"
		ret += " past: " + pastLink(o.title)
	}
	ret += fmt.Sprintf(" self: %s\n", o.selfUrl)
	if o.text != "" {
		ret += fmt.Sprintf("\n%s\n", fixupComment(o.text, 3, o.thread.LineWidth))
	}
	ret += fmt.Sprintf(
		"\n%s(%d) - %s - %d Comments\n",
		format.AuthorStyle.Render(o.user),
		o.score,
		humanize.Time(o.date),
		o.ncomments,
	)
	return
}

func (c Comment) String() (ret string) {
	leftPadding := c.thread.LeftPadding * c.depth
	rightPadding := 2
	lineWidth := format.Min(c.thread.LineWidth, leftPadding+c.thread.CommentWidth+1) - rightPadding
	ret += "\n" + fixupComment(c.msg, leftPadding+1, lineWidth) + "\n"

	arrow := ">> "
	if c.depth > 0 {
		arrow = ">> "
	}

	author := c.user
	if c.user == c.thread.op.user {
		author = format.AuthorStyle.Render(c.user)
	}

	if c.thread.ShowDate {
		ret += strings.Repeat(" ", leftPadding) + arrow + author + " - " + humanize.Time(c.date)
	} else {
		ret += strings.Repeat(" ", leftPadding) + arrow + author
	}

	ret += "\n"
	return
}

var reAnchor = regexp.MustCompile(`<a href="([^"]+)"( rel="nofollow")?>([^<]+)</a>`)
var reItalic = regexp.MustCompile(`<i>([^<]+)</i>`)

var italicStyle = lipgloss.NewStyle().Italic(true)

func fixupComment(htmlText string, leftPad int, width int) string {
	plainText := strings.ReplaceAll(htmlText, "<p>", "\n\n")
	plainText = strings.TrimSpace(plainText)
	plainText = reItalic.ReplaceAllStringFunc(plainText, func(s string) string {
		return italicStyle.Render(s[3 : len(s)-4])
	})
	plainText = reAnchor.ReplaceAllStringFunc(plainText, func(s string) string {
		matches := reAnchor.FindAllStringSubmatch(s, -1)
		href := matches[0][1]
		desc := strings.TrimSuffix(matches[0][3], "...")
		if strings.HasPrefix(href, desc) {
			return href
		} else {
			return fmt.Sprintf("[%s](%s)", desc, href)
		}
	})
	wrapped, _ := text.WrapLeftPadded(
		format.GreenTextIt(html.UnescapeString(plainText)),
		width,
		leftPad,
	)
	return wrapped
}

func pastLink(title string) string {
	return fmt.Sprintf(
		"https://hn.algolia.com/?query=%s&sort=byDate\n",
		url.QueryEscape(title),
	)
}
