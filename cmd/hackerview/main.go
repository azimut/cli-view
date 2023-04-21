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
	timeout   time.Duration
	useColors bool
	useDate   bool
	usePretty bool
	width     uint
	limit     uint
	workers   uint
}

var opt options
var url string

func init() {
	flag.BoolVar(&opt.useDate, "d", false, "print date on comments")
	flag.BoolVar(&opt.useColors, "C", true, "use colors")
	flag.BoolVar(&opt.usePretty, "P", true, "use pretty formatting")
	flag.DurationVar(&opt.timeout, "t", time.Second*5, "timeout in seconds")
	flag.UintVar(&opt.limit, "l", 0, "limits the ammount of comments to fetch")
	flag.UintVar(&opt.width, "w", 100, "fixed with")
	flag.UintVar(&opt.workers, "W", 3, "number of workers to fetch comments")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func run(args []string, stdout io.Writer) error {
	flag.Parse()
	flag.Usage = usage
	color.NoColor = !opt.useColors
	if flag.NArg() != 1 {
		flag.Usage()
		return errors.New("missing URL argument")
	}
	url = flag.Args()[0]
	op, comments, err := hackernews.Fetch(url, opt.timeout, int(opt.limit), opt.workers)
	if err != nil {
		return errors.New("could not fetch url")
	}
	hackernews.Format(int(opt.width), opt.useDate, op, comments)
	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
