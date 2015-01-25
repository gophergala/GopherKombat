package client

import (
	"bytes"
	"encoding/json"
	"github.com/gophergala/GopherKombat/common/request"
	"net/http"
)

const EXECUTION_SERVER = "http://localhost:1212/combat"

func CompareExecution(code1 string, code2 string) (*request.Response, error) {
	req := &request.Request{}
	req.Contestant1.Name = "code1"
	req.Contestant1.Code = code1
	req.Contestant2.Name = "code2"
	req.Contestant2.Code = code2

	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(EXECUTION_SERVER, "application/json", &buff)
	if err != nil {
		return nil, err
	}

	response := &request.Response{}
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
