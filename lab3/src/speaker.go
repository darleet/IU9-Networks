package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func (p *Peer) dial(msg Message) {
	conn, err := net.Dial("tcp", p.NextAddress())
	if err != nil {
		p.logger.Println(err.Error())
		return
	}
	defer conn.Close()

	jBytes, err := json.Marshal(msg)
	if err != nil {
		p.logger.Panic(err.Error())
	}

	_, err = conn.Write(jBytes)
	if err != nil {
		p.logger.Panic(err.Error())
	}
	p.logger.Printf("%s: sent message to %s successfully\n",
		p.Address(), p.NextAddress())
}

func (p *Peer) say() {
	for {
		var cmd string
		fmt.Print("\nEnter one of the commands (add/del/send/stop): ")
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "add":
			var name string
			fmt.Print("Enter name of peer to subscribe: ")
			fmt.Scanf("%s", &name)
			p.mu.Lock()
			p.subscrList[name] = struct{}{}
			p.mu.Unlock()
			fmt.Println("Success!")
		case "del":
			var name string
			fmt.Print("Enter name of peer to unsubscribe: ")
			fmt.Scanf("%s", &name)
			p.mu.Lock()
			delete(p.subscrList, name)
			p.mu.Unlock()
			fmt.Println("Success!")
		case "send":
			var s string
			fmt.Print("Enter message text: ")
			fmt.Scanf("%s", &s)
			go p.dial(Message{p.name, s})
			fmt.Println("Message was sent to next peer (if he is online)")
		case "stop":
			p.stop <- struct{}{}
			return
		default:
			fmt.Println("Incorrect command! Try again.")
		}
	}
}
