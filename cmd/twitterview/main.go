package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/azimut/cli-view/internal/tui"
	"github.com/azimut/cli-view/internal/twitter"
)

type options struct {
	timeout   time.Duration
	useColors bool
	useTUI    bool
	userAgent string
}

var opts options

func init() {
	flag.BoolVar(&opts.useColors, "C", true, "use colors")
	flag.BoolVar(&opts.useTUI, "x", false, "use TUI")
	flag.DurationVar(&opts.timeout, "t", time.Second*5, "timeout in seconds")
	flag.StringVar(&opts.userAgent, "A", "Twitter_View/0.1", "default User-Agent sent")
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

	url := flag.Args()[0]
	tweet, err := twitter.Fetch(url, opts.userAgent, opts.timeout)
	if err != nil {
		return err
	}

	if opts.useTUI {
		tui.RenderLoop(twitter.NewProgram(tweet))
	} else {
		fmt.Println(tweet)
	}

	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
