package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/reddit"
	"github.com/fatih/color"
)

type options struct {
	timeout   time.Duration
	useColors bool
	userAgent string
	width     int
}

var opts options

func init() {
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.BoolVar(&opts.useColors, "C", true, "use colors")
	flag.StringVar(&opts.userAgent, "A", "Reddit_Cli/0.1", "user agent to send to reddit")
	flag.IntVar(&opts.width, "w", 80, "fixed with")
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
	thread, err := reddit.Fetch(url, opts.userAgent, opts.timeout)
	if err != nil {
		return err
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
