package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/lobsters"
	"github.com/azimut/cli-view/internal/tui"
	"github.com/fatih/color"
)

type options struct {
	leftPadding  uint
	timeout      time.Duration
	useColors    bool
	showAuthor   bool
	showDate     bool
	showId       bool
	userAgent    string
	useTUI       bool
	lineWidth    uint
	commentWidth uint
}

var opts options

func init() {
	flag.BoolVar(&opts.useColors, "C", true, "use colors")
	flag.BoolVar(&opts.showDate, "d", false, "print date on comments")
	flag.BoolVar(&opts.showAuthor, "a", true, "print author on comments")
	flag.BoolVar(&opts.useTUI, "x", false, "use TUI")
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opts.userAgent, "A", "LobstersView/1.0", "user agent to send")
	flag.UintVar(&opts.lineWidth, "w", 100, "line width")
	flag.UintVar(&opts.commentWidth, "W", 80, "comment width")
	flag.UintVar(&opts.leftPadding, "l", 3, "left padding on comments")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func run(args []string, stdout io.Writer) error {
	flag.Parse()
	flag.Usage = usage
	color.NoColor = !opts.useColors
	if flag.NArg() != 1 {
		flag.Usage()
		return errors.New("missing URL argument")
	}

	url := flag.Args()[0]
	thread, err := lobsters.Fetch(url, opts.userAgent, opts.timeout)
	if err != nil {
		return errors.New("could not fetch url")
	}
	thread.CommentWidth = int(opts.commentWidth)
	thread.LineWidth = int(opts.lineWidth)
	thread.LeftPadding = int(opts.leftPadding)
	thread.ShowAuthor = opts.showAuthor
	thread.ShowDate = opts.showDate

	if opts.useTUI {
		tui.RenderLoop(lobsters.NewProgram(thread))
	} else {
		fmt.Println(thread)
	}

	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
