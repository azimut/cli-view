package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/twitter"
)

type options struct {
	userAgent string
	timeout   time.Duration
	width     int
}

var opt options
var url string

func init() {
	flag.IntVar(&opt.width, "w", 0, "fixed with, defaults to console width")
	flag.StringVar(&opt.userAgent, "A", "Wget", "default User-Agent sent")
	flag.DurationVar(&opt.timeout, "t", time.Second*5, "timeout in seconds")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	flag.Usage = usage
	url = "https://twitter.com/vickyguareschi/status/1432922904556146691"
	res, err := twitter.Fetch(url, opt.userAgent, opt.timeout)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
