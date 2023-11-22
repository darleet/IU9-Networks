package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"lab5/model"
	"log"
	"os"
	"os/signal"
	"strconv"
)

func listenTerminal(input chan string) {
	var s string
	for {
		fmt.Scan(&s)
		input <- s
	}
}

func handleTerminal(c *websocket.Conn) {
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
				log.Println(fmt.Sprintf("could not convert `%v` to int", s))
				break
			}
			err = sendRequest(c, model.Request{Value: v})
			if err != nil {
				log.Println("sendRequest:", err)
			}
		}
	}
}
