package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()
	r.HandleFunc("/echo", IndexHandler)
	static := "static"
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(static))))

	s := negroni.Classic()
	s.UseHandler(r)
	return s
}
