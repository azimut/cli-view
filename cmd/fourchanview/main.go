package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/azimut/cli-view/internal/fourchan"
	"github.com/fatih/color"
)

type options struct {
	leftPadding uint
	showColors  bool
	width       uint
}

var opts options

func init() {
	flag.BoolVar(&opts.showColors, "C", true, "show colors")
	flag.UintVar(&opts.width, "w", 80, "fixed with")
	flag.UintVar(&opts.leftPadding, "l", 3, "left padding for child comments")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] URL ...\n", os.Args[0])
	flag.PrintDefaults()
}

func run(args []string, stdout io.Writer) error {
	flag.Parse()
	flag.Usage = usage
	color.NoColor = !opts.showColors
	if flag.NArg() != 1 {
		flag.Usage()
		return errors.New("missing URL argument")
	}
	url := flag.Args()[0]
	thread, err := fourchan.Fetch(url, opts.width, opts.leftPadding)
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
