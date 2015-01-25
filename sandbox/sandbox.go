package main

import (
	"encoding/json"
	"fmt"
	"github.com/gophergala/GopherKombat/common/request"
	"log"
	"net/http"
)

func combatHandler(w http.ResponseWriter, r *http.Request) {
	var req request.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("error decoding request: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Printf("%#v\n", req)

	resp := executeCombat(&req)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func executeCombat(req *request.Request) *request.Response {
	resp := &request.Response{}

	engine, ai1Err, ai2Err := NewEngine(req)
	if ai1Err != nil || ai2Err != nil {
		resp.Error1 = fmt.Sprintf("%v", ai1Err)
		resp.Error2 = fmt.Sprintf("%v", ai2Err)
		return resp
	}
	defer engine.Close()

	time1, time2, err1, err2 := engine.Run()
	resp.Time1 = time1
	resp.Time2 = time2
	resp.Error1 = fmt.Sprintf("%v", err1)
	resp.Error2 = fmt.Sprintf("%v", err2)

	return resp
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("root")
	fmt.Fprintf(w, "running")
}

func main() {
	log.Printf("Running")
	http.HandleFunc("/combat", combatHandler)
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":1212", nil))
}
