package reddit

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/azimut/cli-view/internal/format"
	md "github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

var reMarkdownLink = regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`)
var reHTTPLink = regexp.MustCompile(`[^\[^\(^m]http[s]?://[^\s^\[^\(^\[]+`)

func (op Op) String() (ret string) {
	leftPadding := int(op.thread.LeftPadding)
	maxWidth := int(op.thread.Width)
	ret += "\n"
	ret += fmt.Sprintf("title: %s\n", op.title)
	ret += fmt.Sprintf(" self: %s\n", op.self)
	if op.url != op.self && op.url != "" {
		ret += fmt.Sprintf("  url: %s\n", op.url)
	}
	ret += "\n"
	ret += fixupContent(
		op.selftext,
		maxWidth-leftPadding,
		leftPadding,
	)
	ret += "\n"
	ret += fmt.Sprintf("%s(%d) - %s - %d Comment(s)\n\n\n",
		format.AuthorStyle.Render(op.author),
		op.upvotes,
		relativeFromUnix(op.createdUTC),
		op.nComments)
	return
}

func (comment Comment) String() (ret string) {
	leftPadding := int(comment.thread.LeftPadding)
	maxWidth := int(comment.thread.Width)
	commentLeftPadding := int(comment.depth)*leftPadding + 1
	ret += fixupContent(
		comment.message,
		maxWidth,
		commentLeftPadding,
	)

	author := comment.author
	if comment.author == comment.thread.op.author {
		author = format.AuthorStyle.Render(comment.author)
	}

	ret += strings.Repeat(" ", int(comment.depth)*leftPadding)
	ret += fmt.Sprintf(
		">> %s(%d) - %s\n\n",
		author,
		comment.score,
		relativeFromUnix(comment.createdUtc),
	)

	for _, reply := range comment.replies {
		ret += fmt.Sprint(reply)
	}

	return
}

func (post Thread) String() (ret string) {
	ret += fmt.Sprint(post.op)
	for _, comment := range post.comments {
		ret += fmt.Sprintln(comment)
	}
	return
}

func extensions() parser.Extensions {
	extensions := parser.NoIntraEmphasis // Ignore emphasis markers inside words
	extensions |= parser.Tables          // Parse tables
	extensions |= parser.FencedCode      // Parse fenced code blocks
	// extensions |= parser.Autolink               // Detect embedded URLs that are not explicitly marked
	extensions |= parser.Strikethrough          // Strikethrough text using ~~test~~
	extensions |= parser.SpaceHeadings          // Be strict about prefix heading rules
	extensions |= parser.HeadingIDs             // specify heading IDs  with {#id}
	extensions |= parser.BackslashLineBreak     // Translate trailing backslashes into line breaks
	extensions |= parser.DefinitionLists        // Parse definition lists
	extensions |= parser.LaxHTMLBlocks          // more in HTMLBlock, less in HTMLSpan
	extensions |= parser.NoEmptyLineBeforeBlock // no need for new line before a list
	return extensions
}

func fixupContent(content string, width, padding int) (ret string) {
	ret = reMarkdownLink.ReplaceAllStringFunc(content, func(s string) string {
		matches := reMarkdownLink.FindAllStringSubmatch(s, -1)
		if matches[0][1] == matches[0][2] {
			// color.New(color.FgBlue).Sprint(matches[0][1])
			// fmt.Sprintf("[_](%s)", matches[0][1])
			return matches[0][1]
		} else {
			return s
		}
	})

	ret = strings.Replace(ret, "&amp;#x200B;", "", -1)
	ret = strings.Replace(ret, "&gt;", ">", -1)
	ret = strings.Replace(ret, "&lt;", "<", -1)

	// ret = string(markdown.Render(ret, width, padding))
	p := parser.NewWithExtensions(extensions())
	nodes := md.Parse([]byte(ret), p)
	renderer := markdown.NewRenderer(width, padding)
	ret = string(md.Render(nodes, renderer))

	ret = reHTTPLink.ReplaceAllStringFunc(ret, func(s string) string {
		return color.New(color.FgBlue).Sprint(s)
	})
	return
}

func relativeFromUnix(unix int64) string {
	return humanize.Time(time.Unix(unix, 0))
}
