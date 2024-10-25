package main

// ignore

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/itchio/dash"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTIONS] search-path\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	// mainly for dash logging
	flag.Usage = usage
	flag.BoolVar(&settings.debug, "debug", false, "show debug messages")
	flag.BoolVar(&settings.quiet, "quiet", false, "only show errors")
	flag.Parse()

	searchPath := flag.Arg(0)
	if searchPath == "" {
		fmt.Fprint(os.Stderr, "error: Missing search path\n\n")
		usage()
		os.Exit(1)
	}

	// send to dash...
	verdict, err := dash.Configure(searchPath, dash.ConfigureParams{
		Consumer: setupState(),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't find candidate: %s\n", err)
		os.Exit(1)
	}
	for _, c := range verdict.Candidates {
		// don't bother printing libmagic metadata
		c.Spell = nil
	}

	// ... and serialize it
	res, err := json.Marshal(verdict)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create json: %s\n", err)
		os.Exit(1)
	}
	fmt.Print(string(res))
}
