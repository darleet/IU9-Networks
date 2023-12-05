package service

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
		p.Logger.Panic(err.Error())
	}

	if msg.PeerName == p.Name {
		return
	}
	go p.dial(Message{msg.PeerName, msg.Text})

	if _, ok := p.SubscrList[msg.PeerName]; ok {
		fmt.Printf("\nMessage from %s: %s\n", msg.PeerName, msg.Text)
		p.Logger.Printf("%s: got message from %s and read it\n",
			p.Address(), p.NextAddress())
	} else {
		p.Logger.Printf("%s: got message from %s and didnt'r read it\n",
			p.Address(), p.NextAddress())
	}
}

func (p *Peer) listen() {
	l, err := net.Listen("tcp", p.Address())
	if err != nil {
		p.Logger.Println(err.Error())
		p.Stop <- struct{}{}
		panic(err.Error())
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			select {
			case <-p.Stop:
				return
			default:
				p.Logger.Panic(err.Error())
			}
		} else {
			go p.handleConnection(conn)
		}
	}
}
