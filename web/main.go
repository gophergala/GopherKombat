package main

import (
	"github.com/gophergala/GopherKombat/web/login"
	"net/http"
)

const (
	FILE_SERVE_PATH = "/var/static"
)

func main() {
	StartWebServer()
}

func StartWebServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(FILE_SERVE_PATH)))
	mux.HandleFunc("/login/callback", login.LoginCallback)

	panic(http.ListenAndServe(":8080", mux))
}
