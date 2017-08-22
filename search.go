package main

import (
	// "fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

const version = "0.9"

type Options struct {
	Recursion bool `short:"r" description:"Recursive search"`
}

var opts Options

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "search"
	parser.Usage = "[OPTIONS] PATTERN [PATH]"

	args, _ := parser.Parse()

	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
