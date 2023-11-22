package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func Run(address string) {
	u := url.URL{Scheme: "ws", Host: address, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go listenSocket(c)
	handleTerminal(c)
}
