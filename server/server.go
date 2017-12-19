package server

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func GetServer() *negroni.Negroni{
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	s := negroni.Classic()
	s.UseHandler(r)
	return s
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not upgrade to websocket")
		os.Exit(1)
	}
	defer ws.Close()

	for {
		mType, b, err := ws.ReadMessage()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can not read message.")
			os.Exit(1)
		}
		err = ws.WriteMessage(mType, b)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can not write message")
			os.Exit(1)
		}
	}
}