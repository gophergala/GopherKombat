package request

import (
	"time"
)

type Contestant struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Request struct {
	Contestant1 Contestant `json:"ai1"`
	Contestant2 Contestant `json:"ai2"`
}

type Response struct {
	Time1  time.Duration `json:"t1"`
	Time2  time.Duration `json:"t2"`
	Error1 string        `json:"err1"`
	Error2 string        `json:"err2"`
}
