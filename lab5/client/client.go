package client

import (
	"log"
	"net/url"
	"websockets/client/service"

	"github.com/gorilla/websocket"
)

func Run(address string) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go service.ListenSocket(c)
	service.HandleTerminal(c)
}
