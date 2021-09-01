package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/tui"
	"github.com/azimut/cli-view/internal/twitter"
)

type options struct {
	timeout   time.Duration
	userAgent string
	width     int
}

var opt options
var url string

func init() {
	flag.DurationVar(&opt.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opt.userAgent, "A", "Wget", "default User-Agent sent")
	flag.IntVar(&opt.width, "w", 0, "fixed with, defaults to console width")
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
		fmt.Println("could not fetch url:", err)
		os.Exit(1)
	}
	p := tui.NewProgram(res.Html)
	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
