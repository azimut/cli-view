package format

import (
	"math"
	"strings"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)

func GreenTextIt(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, ">") {
			lines[i] = green.Sprint(line)
		}
	}
	return strings.Join(lines, "\n")
}

func WrapLeftPadded(text string, maxWidth, padding int) string {
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
