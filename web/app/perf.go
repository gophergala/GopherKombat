package app

import (
	"github.com/gophergala/GopherKombat/sandbox/client"
	"log"
	"net/http"
)

func PerfHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := GetCurrentUser(r)
	data := make(map[string]interface{})
	data["loggedIn"] = ok
	if ok {
		data["user"] = user
	}
	render(w, "perf", data)

}

func PerfSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, ok := GetCurrentUser(r)
		data := make(map[string]interface{})
		data["success"] = ok
		if ok {
			r.ParseForm()
			code1 := r.PostFormValue("code1")
			code2 := r.PostFormValue("code2")
			res, err := client.CompareExecution(code1, code2)
			log.Println(res)
			if err != nil || res.Error1 != "" || res.Error2 != "" {
				data["success"] = false
			}
			data["resp"] = res
		} else {
			data["message"] = "You are not logged in."
		}
		renderJson(w, r, data)
	} else {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
	}
}
