package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophergala/GopherKombat/common/game"
	"log"
	"net/http"
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
}

func combatHandler(w http.ResponseWriter, r *http.Request) {
	var req Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Printf("%#v\n", req)

	resp, err := executeCombat(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func executeCombat(req *Request) (*Response, error) {
	resp := &Response{}

	cp, err := NewContestantProcess(&req.Contestant1)
	if err != nil {
		return nil, err
	}
	defer cp.Close()

	state := &game.State{Test: "test state"}
	action, err := cp.Turn(state)
	if err != nil {
		return nil, err
	}
	log.Printf("%#v", action)
	action, err = cp.Turn(state)
	if err != nil {
		return nil, err
	}
	log.Printf("%#v", action)

	return resp, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("root")
	fmt.Fprintf(w, "running")
}

func main() {
	log.Printf("Running")
	http.HandleFunc("/combat", combatHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
