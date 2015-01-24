package app

import (
	"github.com/gorilla/sessions"
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

func GetCurrentUser(r *http.Request) (string, bool) {
	session := InitSession(r)
	if session.Values["user"] == nil {
		return "", false
	} else {
		return session.Values["user"].(string), true
	}
}
