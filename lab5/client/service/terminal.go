package service

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"websockets/model"

	"github.com/gorilla/websocket"
)

func listenTerminal(input chan string) {
	var s string
	for {
		fmt.Scan(&s)
		input <- s
	}
}

func HandleTerminal(c *websocket.Conn) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string)
	go listenTerminal(input)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			err := closeSocket(c)
			if err != nil {
				log.Panic("closeSocket:", err)
			}
			return
		case s := <-input:
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Printf("could not convert `%v` to int\n", s)
				break
			}
			err = sendRequest(c, model.Request{Value: v})
			if err != nil {
				log.Println("sendRequest:", err)
			}
		}
	}
}
