package main

import (
	"log"
)

const (
	MAP_WIDTH  = 20
	MAP_HEIGHT = 20
)

type Tile struct {
}

type Engine struct {
	ai1 *ContestantProcess
	ai2 *ContestantProcess

	board [][]Tile
}

func NewEngine(request *Request) (*Engine, error, error) {
	var ai1Err, ai2Err error
	engine := &Engine{}
	engine.ai1, ai1Err = NewContestantProcess(&request.Contestant1)
	engine.ai2, ai2Err = NewContestantProcess(&request.Contestant2)
	if ai1Err != nil || ai2Err != nil {
		return nil, ai1Err, ai2Err
	}

	// initialize board
	engine.board = make([][]Tile, MAP_HEIGHT)
	for i := 0; i < MAP_HEIGHT; i++ {
		engine.board[i] = make([]Tile, MAP_WIDTH)
	}

	return engine, nil, nil
}

func (eng *Engine) Run() error {
	log.Printf("running combat")
	_, err := eng.ai1.Run()
	if err != nil {
		return err
	}
	_, err = eng.ai2.Run()
	if err != nil {
		return err
	}

	return nil
}

func (eng *Engine) Close() {
	eng.ai1.Close()
	eng.ai2.Close()
}
