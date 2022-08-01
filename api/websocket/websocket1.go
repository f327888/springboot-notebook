package main

import "fmt"
import websocket "code.google.com/p/go/websocket"

func server(ws *websocket.Conn) {
	fmt.Printf("new connettion\n")
	buf := make([]byte, 100)
	for {
		ws.Read
	}
}

func main() {

}
