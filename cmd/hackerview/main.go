package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/hackernews"
	"github.com/azimut/cli-view/internal/tui"
	"github.com/fatih/color"
)

type options struct {
	leftPadding uint
	maxComments uint
	nWorkers    uint
	timeout     time.Duration
	showColors  bool
	showDate    bool
	useTUI      bool
	width       uint
}

var opts options

func init() {
	flag.BoolVar(&opts.showColors, "C", true, "use colors")
	flag.BoolVar(&opts.showDate, "d", false, "print date on comments")
	flag.BoolVar(&opts.useTUI, "x", false, "use TUI")
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.UintVar(&opts.maxComments, "c", 10, "limits the ammount of comments to fetch")
	flag.UintVar(&opts.nWorkers, "W", 3, "number of workers to fetch comments")
	flag.UintVar(&opts.width, "w", 100, "fixed with")
	flag.UintVar(&opts.leftPadding, "l", 3, "left padding")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func run(args []string, stdout io.Writer) error {
	flag.Parse()
	flag.Usage = usage
	color.NoColor = !opts.showColors
	if flag.NArg() != 1 {
		flag.Usage()
		return errors.New("missing URL argument")
	}

	url := flag.Args()[0]
	thread, err := hackernews.Fetch(url, opts.timeout, opts.maxComments, opts.nWorkers)
	if err != nil {
		return errors.New("could not fetch url")
	}
	thread.Width = opts.width
	thread.LeftPadding = opts.leftPadding
	thread.ShowDate = opts.showDate

	if opts.useTUI {
		tui.RenderLoop(hackernews.NewProgram(thread))
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
