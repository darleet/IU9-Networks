package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func (p *Peer) handleConnection(c net.Conn) {
	defer c.Close()
	d := json.NewDecoder(c)

	var msg Message
	err := d.Decode(&msg)
	if err != nil {
		p.logger.Panic(err.Error())
	}

	if msg.PeerName == p.name {
		return
	}
	go p.dial(Message{msg.PeerName, msg.Text})

	if _, ok := p.subscrList[msg.PeerName]; ok {
		fmt.Printf("\nMessage from %s: %s\n", msg.PeerName, msg.Text)
		p.logger.Printf("%s: got message from %s and read it\n",
			p.Address(), p.NextAddress())
	} else {
		p.logger.Printf("%s: got message from %s and didnt'r read it\n",
			p.Address(), p.NextAddress())
	}
}

func (p *Peer) listen() {
	l, err := net.Listen("tcp", p.Address())
	if err != nil {
		p.logger.Println(err.Error())
		p.stop <- struct{}{}
		panic(err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			select {
			case <-p.stop:
				return
			default:
				p.logger.Panic(err.Error())
			}
		} else {
			go p.handleConnection(conn)
		}
	}
}
