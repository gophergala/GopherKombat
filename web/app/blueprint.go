package app

import (
	"io/ioutil"
	"log"
	"net/http"
)

func BlueprintHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := GetCurrentUser(r)
	if ok {
		data, err := ioutil.ReadFile(FILE_SERVE_PATH + "/template/blueprint.html")
		if err != nil {
			log.Printf("Error reading template: %s ", err)
		}
		w.Write(data)
	}

}
