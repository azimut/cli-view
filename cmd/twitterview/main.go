package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/twitter"
)

type options struct {
	timeout   time.Duration
	userAgent string
	useColors bool
	width     int
}

var opt options
var url string

func init() {
	flag.DurationVar(&opt.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opt.userAgent, "A", "Wget", "default User-Agent sent")
	flag.BoolVar(&opt.useColors, "C", true, "use colors")
	flag.IntVar(&opt.width, "w", 0, "fixed with, defaults to console width")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	flag.Usage = usage
	if flag.NArg() != 1 {
		flag.Usage()
		panic("Error: Missing URL argument.")
	}
	url = flag.Args()[0]
	res, err := twitter.Fetch(url, opt.userAgent, opt.timeout)
	if err != nil {
		fmt.Println("could not fetch url:", err)
		os.Exit(1)
	}
	fmt.Println(twitter.Format(res))
}
