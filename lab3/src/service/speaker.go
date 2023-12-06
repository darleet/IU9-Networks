package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func (p *Peer) dial(msg Message) {
	conn, err := net.Dial("tcp", p.NextAddress())
	if err != nil {
		p.Logger.Println(err.Error())
		return
	}
	defer conn.Close()

	jBytes, err := json.Marshal(msg)
	if err != nil {
		p.Logger.Panic(err.Error())
	}

	_, err = conn.Write(jBytes)
	if err != nil {
		p.Logger.Panic(err.Error())
	}
	p.Logger.Printf("%s: sent message to %s successfully\n",
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
			p.SubscrList[name] = struct{}{}
			p.mu.Unlock()
			fmt.Println("Success!")
		case "del":
			var name string
			fmt.Print("Enter name of peer to unsubscribe: ")
			fmt.Scanf("%s", &name)
			p.mu.Lock()
			delete(p.SubscrList, name)
			p.mu.Unlock()
			fmt.Println("Success!")
		case "send":
			fmt.Print("Enter message text: ")
			r := bufio.NewReader(os.Stdin)
			s, err := r.ReadString('\n')
			if err != nil {
				fmt.Println("Error occurred while reading message text")
			} else {
				go p.dial(Message{p.Name, s})
				fmt.Println("Message was sent to next peer (if he is online)")
			}
		case "stop":
			p.Stop <- struct{}{}
			return
		default:
			fmt.Println("Incorrect command! Try again.")
		}
	}
}
