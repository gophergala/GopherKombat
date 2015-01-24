package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

func RankingsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(FILE_SERVE_PATH + "/template/rankings.html")
	if err != nil {
		log.Printf("Error reading template: %s ", err)
	}
	w.Write(data)
}
