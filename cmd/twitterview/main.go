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

func init() {
	flag.StringVar(&opt.userAgent, "A", "CliView/0.1", "default User-Agent sent")
	flag.DurationVar(&opt.timeout, "t", time.Second*5, "timeout in seconds")
	flag.IntVar(&opt.width, "w", 0, "fixed with, defaults to console width")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	flag.Usage = usage
	t := twitter.Tweet{}
	t.GetHeader("what is this")
	println("HELP ME")
}
