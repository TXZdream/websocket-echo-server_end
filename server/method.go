package server

import (
	"fmt"
	"net/url"
	"os"

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
func HandleMessage(mType int, msg []byte, conn *websocket.Conn, isFirst bool) {
	var workSpace *Command
	var err error
	if isFirst {
		pwd := GetPwd("test")
		var env []string
		entrypoint := make([]string, 1) // Set `/go` as default entrypoint
		entrypoint[0] = "/go"
		username := "test"
		workSpace = &Command{
			Command:    string(msg),
			Entrypoint: entrypoint,
			PWD:        pwd,
			ENV:        env,
			UserName:   username,
		}
	}

	// Send message
	if isFirst {
		err = conn.WriteJSON(*workSpace)
	} else {
		err = conn.WriteMessage(mType, msg)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not write message to connection")
		return
	}
}

// InitDockerConnection inits the connection to the docker service with the first message received from client
func InitDockerConnection(msg string) *websocket.Conn {
	// Just handle command start with `go`
	if len(msg) > 3 && msg[0:3] == "go " {
		conn := DialDockerService()
		if conn == nil {
			return nil
		}
		return conn
	}
	return nil
}

// GetPwd return current path of given username
func GetPwd(username string) string {
	// Return user root in test version
	return "/"
}
