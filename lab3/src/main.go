package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Message struct {
	PeerName string `json:"peerName"`
	Text     string `json:"text"`
}

type Peer struct {
	name       string
	peerIP     string
	port       string
	nextPeerIP string
	nextPort   string
	logger     *log.Logger
	stop       chan struct{}
	mu         sync.Mutex
	subscrList map[string]struct{}
}

func (p *Peer) Address() string {
	return p.peerIP + ":" + p.port
}

func (p *Peer) NextAddress() string {
	return p.nextPeerIP + ":" + p.nextPort
}

func (p *Peer) StartPeer() {
	go p.listen()
	go p.say()
	<-p.stop
	p.logger.Printf("%s: peer shutdown\n", p.Address())
	fmt.Println("Peer shutting down...")
}

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

	p := &Peer{
		name:       name,
		peerIP:     peerIP,
		port:       port,
		nextPeerIP: nextPeerIP,
		nextPort:   nextPort,
		logger:     logger,
		stop:       make(chan struct{}),
		subscrList: make(map[string]struct{}),
	}

	p.StartPeer()
}
