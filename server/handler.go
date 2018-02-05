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
	// Set TextMessage as default
	msgType := websocket.TextMessage
	clientMsg := make(chan []byte)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not upgrade to websocket")
		os.Exit(1)
	}
	defer ws.Close()

	// Open a goroutine to receive message from client connection
	go func(clientChan chan<- []byte, conn *websocket.Conn) {
		for {
			_, b, err := ws.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseGoingAway) {
					fmt.Fprintln(os.Stderr, "Remote user closed the connection")
					ws.Close()
					close(clientChan)
					break
				}
				close(clientChan)
				fmt.Fprintln(os.Stderr, "Can not read message.")
				return
			}
			clientChan <- b
		}
	}(clientMsg, ws)

	// Handle messages from the channel
	isFirst := true
	for msg := range clientMsg {
		var conn *websocket.Conn
		// Init the connection to the docker serveice
		if isFirst {
			conn = InitDockerConnection(string(msg))
			if conn == nil {
				fmt.Fprintf(os.Stderr, "Invalid command.")
				ws.WriteMessage(msgType, []byte("Invalid Command"))
				return
			}
			// Listen message from docker service and send to client connection
			go func(cConn *websocket.Conn, sConn *websocket.Conn) {
				defer sConn.Close()
				for {
					mType, msg, err := sConn.ReadMessage()
					if err != nil {
						fmt.Fprintln(os.Stderr, "Can not read message from connection")
						return
					}
					cConn.WriteMessage(mType, msg)
				}
			}(ws, conn)
		}

		// Send message to docker service
		HandleMessage(msgType, msg, conn, isFirst)
		isFirst = false
	}
	ws.Close()
}
