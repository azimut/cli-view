package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/twitter"
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
	flag.IntVar(&opt.width, "w", 0, "fixed with, defaults to console width")
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
	res, err := twitter.Fetch(url, opt.userAgent, opt.timeout)
	if err != nil {
		return errors.New("could not fetch url")
	}
	fmt.Println(twitter.Format(res))
	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
