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
)

type options struct {
	timeout   time.Duration
	userAgent string
	useColors bool
	usePretty bool
	width     int
}

var opt options
var url string

func init() {
	flag.DurationVar(&opt.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opt.userAgent, "A", "Wget", "default User-Agent sent")
	flag.BoolVar(&opt.useColors, "C", true, "use colors")
	flag.BoolVar(&opt.usePretty, "P", true, "use pretty formatting")
	flag.IntVar(&opt.width, "w", 80, "fixed with")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func run(args []string, stdout io.Writer) error {
	flag.Parse()
	flag.Usage = usage
	if flag.NArg() != 1 {
		flag.Usage()
		return errors.New("missing URL argument")
	}
	url = flag.Args()[0]
	op, comments, err := hackernews.Fetch(url, opt.userAgent, opt.timeout)
	if err != nil {
		return errors.New("could not fetch url")
	}
	hackernews.Format(opt.width, op, comments)
	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
