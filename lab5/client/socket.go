package client

import (
	"github.com/gorilla/websocket"
	"lab5/model"
	"log"
)

func listenSocket(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("listenSocket:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func closeSocket(c *websocket.Conn) error {
	err := c.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	return err
}

func sendRequest(c *websocket.Conn, msg model.Request) error {
	err := c.WriteJSON(msg)
	return err
}
