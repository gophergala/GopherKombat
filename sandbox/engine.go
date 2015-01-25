package main

import (
	"github.com/gophergala/GopherKombat/common/request"
	"log"
)

type Engine struct {
	ai1 *ContestantProcess
	ai2 *ContestantProcess
}

func NewEngine(request *request.Request) (*Engine, error, error) {
	var ai1Err, ai2Err error
	engine := &Engine{}
	engine.ai1, ai1Err = NewContestantProcess(&request.Contestant1)
	engine.ai2, ai2Err = NewContestantProcess(&request.Contestant2)
	if ai1Err != nil || ai2Err != nil {
		return nil, ai1Err, ai2Err
	}

	return engine, nil, nil
}

func (eng *Engine) Run() (*ProcessResult, *ProcessResult, error, error) {
	log.Printf("running time comparison")
	res1, err1 := eng.ai1.Run()
	res2, err2 := eng.ai2.Run()
	return res1, res2, err1, err2
}

func (eng *Engine) Close() {
	eng.ai1.Close()
	eng.ai2.Close()
}
