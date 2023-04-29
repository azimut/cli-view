package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/reddit"
	"github.com/azimut/cli-view/internal/tui"
	"github.com/fatih/color"
)

type options struct {
	leftPadding  uint
	lineWidth    uint
	commentWidth uint
	timeout      time.Duration
	useColors    bool
	userAgent    string
	useTUI       bool
}

var opts options

func init() {
	flag.BoolVar(&opts.useTUI, "x", false, "use TUI")
	flag.BoolVar(&opts.useColors, "C", true, "use colors")
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opts.userAgent, "A", "Reddit_Cli/0.1", "user agent to send to reddit")
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
	thread, err := reddit.Fetch(url, opts.userAgent, opts.timeout)
	if err != nil {
		return err
	}
	thread.CommentWidth = int(opts.commentWidth)
	thread.LineWidth = int(opts.lineWidth)
	thread.LeftPadding = int(opts.leftPadding)

	if opts.useTUI {
		tui.RenderLoop(reddit.NewProgram(thread))
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
