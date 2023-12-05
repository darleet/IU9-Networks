package service

import (
	"fmt"
	"log"
	"sync"
)

type Message struct {
	PeerName string `json:"peerName"`
	Text     string `json:"text"`
}

type Peer struct {
	Name       string
	IP         string
	Port       string
	NextIP     string
	NextPort   string
	Logger     *log.Logger
	Stop       chan struct{}
	mu         sync.Mutex
	SubscrList map[string]struct{}
}

func (p *Peer) Address() string {
	return p.IP + ":" + p.Port
}

func (p *Peer) NextAddress() string {
	return p.NextIP + ":" + p.NextPort
}

func (p *Peer) StartPeer() {
	go p.listen()
	go p.say()
	<-p.Stop
	p.Logger.Printf("%s: peer shutdown\n", p.Address())
	fmt.Println("Peer shutting down...")
}
