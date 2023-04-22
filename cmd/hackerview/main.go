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
	"github.com/fatih/color"
)

type options struct {
	maxComments uint
	nWorkers    uint
	timeout     time.Duration
	useColors   bool
	useDate     bool
	usePretty   bool
	width       uint
}

var opts options

func init() {
	flag.BoolVar(&opts.useColors, "C", true, "use colors")
	flag.BoolVar(&opts.useDate, "d", false, "print date on comments")
	flag.BoolVar(&opts.usePretty, "P", true, "use pretty formatting")
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.UintVar(&opts.maxComments, "l", 10, "limits the ammount of comments to fetch")
	flag.UintVar(&opts.nWorkers, "W", 3, "number of workers to fetch comments")
	flag.UintVar(&opts.width, "w", 100, "fixed with")
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
	op, comments, err := hackernews.Fetch(url, opts.timeout, opts.maxComments, opts.nWorkers)
	if err != nil {
		return errors.New("could not fetch url")
	}
	hackernews.Format(int(opts.width), opts.useDate, op, comments)
	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
