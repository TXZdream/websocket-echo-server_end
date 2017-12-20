package main
import (
	"os"
	flag "github.com/spf13/pflag"
	"github.com/txzdream/websocket-echo-server_end/server"
)

func main() {
	var PORT = os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = "8080"
	}
	var port = flag.StringP("port", "p", PORT, "Define the port where server runs")
	flag.Parse()

	s := server.GetServer()
	s.Run(":" + *port)
}