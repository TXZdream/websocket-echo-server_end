package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// IndexHandler handle the index(/) uri with websocket
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
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				fmt.Fprintln(os.Stderr, "Remote user closed the connection")
				ws.Close()
				break
			}
			fmt.Println(err)
			fmt.Fprintln(os.Stderr, "Can not read message.")
			os.Exit(1)
		}

		msg := string(b)
		msgChannel := make(chan []byte, 10)
		go HandleMessage(msg, msgChannel)
		for msg := range msgChannel {
			ws.WriteMessage(mType, msg)
		}
		// err = ws.WriteMessage(mType, b)
		// if err != nil {
		// 	if websocket.IsCloseError(err, websocket.CloseGoingAway) {
		// 		fmt.Fprintln(os.Stderr, "Remote user closed the connection")
		// 		ws.Close()
		// 		break
		// 	}
		// 	fmt.Fprintln(os.Stderr, "Can not write message")
		// 	os.Exit(1)
		// }
	}
}
