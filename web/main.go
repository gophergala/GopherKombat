package main

import (
	"encoding/gob"
	"github.com/gophergala/GopherKombat/common/user"
	"github.com/gophergala/GopherKombat/web/app"
	"github.com/gophergala/GopherKombat/web/login"
	"github.com/gorilla/context"
	"net/http"
)

func main() {
	StartWebServer()
}

func StartWebServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(app.FILE_SERVE_PATH)))
	mux.HandleFunc("/login/callback", login.LoginCallbackHandler)
	mux.HandleFunc("/logout", login.LogoutHandler)
	mux.HandleFunc("/blueprint", app.BlueprintHandler)
	mux.HandleFunc("/kombat", app.KombatHandler)
	mux.HandleFunc("/rankings", app.RankingsHandler)
	gob.Register(&user.User{})
	panic(http.ListenAndServe(":8080", context.ClearHandler(mux)))
}
