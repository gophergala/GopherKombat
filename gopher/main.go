package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const TEAM_SIZE = 5

type State struct {
	GopherId int
	Test     string
}

type Action struct {
	Test string
}

func main() {
	gophers := make([]Gopher, TEAM_SIZE)

	for i, gopher := range gophers {
		gopher.Id = i
		gopher.Init()
	}

	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	var state State

	for {
		// Read state from input
		if err := decoder.Decode(&state); err == io.EOF {
			break
		} else if err != nil {
			log.Printf("error reading input: %v", err)
		}

		// Process turn
		action := gophers[state.GopherId].Turn(&state)

		// Write action to stdout
		if err := encoder.Encode(action); err != nil {
			log.Printf("error writing output: %v", err)
		}
	}

	log.Printf("finished running AI")
}
