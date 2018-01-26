package server

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

// DialDockerService create connection between web server and docker server
func DialDockerService() *websocket.Conn {
	// Set up websocket connection
	dockerAddr := "localhost:8081"
	url := url.URL{Scheme: "ws", Host: dockerAddr, Path: "/"}
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not dial docker websocket service")
		return nil
	}
	return conn
}

// HandleMessage decide different operation according to the given json message
func HandleMessage(msg string, msgChannel chan<- []byte) {
	defer close(msgChannel)
	// Just handle command start with `go`
	if len(msg) > 3 && msg[0:3] == "go " {
		pwd := GetPwd("test")
		var env []string
		// Set `/go` as default entrypoint
		entrypoint := make([]string, 1)
		entrypoint[0] = "/go"
		username := "test"
		workSpace := &Command{
			Command:    msg,
			Entrypoint: entrypoint,
			PWD:        pwd,
			ENV:        env,
			UserName:   username,
		}
		conn := DialDockerService()
		if conn == nil {
			return
		}
		defer conn.Close()

		// Read message from connection
		var wg sync.WaitGroup
		wg.Add(1)
		go func(c chan<- []byte, conn *websocket.Conn) {
			defer conn.Close()
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Can not read message from connection")
					wg.Done()
				}
				c <- msg
			}
		}(msgChannel, conn)

		// Send message
		err := conn.WriteJSON(*workSpace)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Can not write message to connection")
			return
		}
		wg.Wait()
	}
}

// GetPwd return current path of given username
func GetPwd(username string) string {
	// Return user root in test version
	return "/"
}
