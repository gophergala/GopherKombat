package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

func KombatHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(FILE_SERVE_PATH + "/template/kombat.html")
	if err != nil {
		log.Printf("Error reading template: %s ", err)
	}
	w.Write(data)
}
