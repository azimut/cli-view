package format

import (
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"github.com/jaytaylor/html2text"
)

var green = color.New(color.FgGreen)
var AuthorStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("8")).
	Foreground(lipgloss.Color("0"))

// GreenTextIt makes green the lines in `text` that start with ">"
func GreenTextIt(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, ">") {
			lines[i] = green.Sprint(line)
		}
	}
	return strings.Join(lines, "\n")
}

// FormatHtml2Text converts an html text into a plaintext one
func FormatHtml2Text(htmlText string, width, leftPadding int) string {
	plainText, err := html2text.FromString(htmlText, html2text.Options{})
	if err != nil {
		panic(err)
	}
	wrapped, _ := text.WrapLeftPadded(GreenTextIt(plainText), width, leftPadding)
	return wrapped
}

func FormatText(plainText string, width, leftPadding int) string {
	wrapped, _ := text.WrapLeftPadded(
		GreenTextIt(strings.ReplaceAll(plainText, "\r\n", "\n")),
		width,
		leftPadding,
	)
	return wrapped
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
