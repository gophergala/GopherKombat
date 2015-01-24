package app

import (
	"github.com/gophergala/GopherKombat/common/user"
	"github.com/gorilla/sessions"
	"html/template"
	"io"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

const (
	FILE_SERVE_PATH = "/var/static"
)

func InitSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "gopher-kombat")
	if err != nil {
		println("ovde")
	}
	return session
}

func GetCurrentUser(r *http.Request) (*user.User, bool) {
	session := InitSession(r)
	if session.Values["user"] == nil {
		return new(user.User), false
	} else {
		return session.Values["user"].(*user.User), true
	}
}

func render(w io.Writer, name string, data interface{}) {
	//data, err := ioutil.ReadFile(FILE_SERVE_PATH + "/template/" + name + ".html")
	//	if err != nil {
	//		log.Printf("Error reading template: %s ", err)
	//	}
	//	w.Write(data)
	t, err := template.ParseFiles(FILE_SERVE_PATH + "/template/" + name + ".html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, data)
}
