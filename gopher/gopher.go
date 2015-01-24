package main

import (
	"log"
)

type Gopher struct {
	Id int
}

func (gopher *Gopher) Init() {
	log.Printf("Init gopher %d\n", gopher.Id)
}

func (gopher *Gopher) Turn(state *State) *Action {
	log.Printf("Turn: %#v\n", state)
	return &Action{Test: "Test action"}
}
