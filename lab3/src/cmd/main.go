package main

import (
	"fmt"
	"log"
	"os"
	"peer/src/service"
)

func main() {
	var name, peerIP, port, nextPeerIP, nextPort string

	fmt.Print("Enter peer name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter peer IP: ")
	fmt.Scanf("%s", &peerIP)
	fmt.Print("Enter peer port: ")
	fmt.Scanf("%s", &port)
	fmt.Print("Enter next peer IP: ")
	fmt.Scanf("%s", &nextPeerIP)
	fmt.Print("Enter next peer port: ")
	fmt.Scanf("%s", &nextPort)

	logger := log.Default()
	f, err := os.OpenFile("logFile.log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	logger.SetOutput(f)

	p := &service.Peer{
		Name:       name,
		IP:         peerIP,
		Port:       port,
		NextIP:     nextPeerIP,
		NextPort:   nextPort,
		Logger:     logger,
		Stop:       make(chan struct{}),
		SubscrList: make(map[string]struct{}),
	}

	p.StartPeer()
}
