package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/vichan"
	"github.com/fatih/color"
)

type options struct {
	leftPadding uint
	timeout     time.Duration
	useColors   bool
	showAuthor  bool
	showDate    bool
	showId      bool
	userAgent   string
	width       uint
}

var opts options

func init() {
	flag.BoolVar(&opts.useColors, "C", true, "use colors")
	flag.BoolVar(&opts.showDate, "d", false, "print date on comments")
	flag.BoolVar(&opts.showAuthor, "a", true, "print author on comments")
	flag.BoolVar(&opts.showId, "i", true, "print id on comments")
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opts.userAgent, "A", "VichanView/1.0", "user agent to send")
	flag.UintVar(&opts.width, "w", 100, "terminal width")
	flag.UintVar(&opts.leftPadding, "l", 3, "left padding")
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
	thread, err := vichan.Fetch(
		url,
		opts.userAgent,
		opts.width,
		opts.leftPadding,
		opts.timeout,
		opts.showAuthor,
		opts.showDate,
		opts.showId,
	)
	if err != nil {
		return errors.New("could not fetch url")
	}
	fmt.Println(thread)
	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}