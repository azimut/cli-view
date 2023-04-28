package format

import (
	"strings"

	text "github.com/MichaelMure/go-term-text"
	"github.com/fatih/color"
	"github.com/jaytaylor/html2text"
)

var green = color.New(color.FgGreen)

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
