package main

import (
	"log"
	"os"

	"github.com/itchio/headway/state"
)

var settings = &struct {
	quiet bool
	debug bool
}{
	false,
	false,
}

func setupState() *state.Consumer {
	log.SetOutput(os.Stderr)

	return &state.Consumer{
		OnMessage: func(level, msg string) {
			if level == "debug" && !settings.debug {
				return
			}
			if level != "error" && settings.quiet {
				return
			}
			log.Printf("[%s] %s", level, msg)
		},
	}
}
