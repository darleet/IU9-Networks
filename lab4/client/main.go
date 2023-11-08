package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

type client struct {
	config     *ssh.ClientConfig
	connection *ssh.Client
	logger     *slog.Logger
}

func (c *client) start() {
	fmt.Printf("Connected to %s\n", c.connection.RemoteAddr())
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		cmd := scanner.Text()

		if strings.Compare(cmd, "exit") == 0 {
			return
		}

		s, err := c.connection.NewSession()
		if err != nil {
			c.logger.Error(err.Error())
			return
		}

		s.Stdout = os.Stdout
		s.Stderr = os.Stderr
		err = s.Run(cmd)
		if err != nil {
			fmt.Println("command execution failed")
			c.logger.Error(err.Error())
		}
		s.Close()
	}
}

func main() {
	lg := slog.Default()

	env, err := godotenv.Read()
	if err != nil {
		lg.Error(err.Error())
		return
	}

	conf := &ssh.ClientConfig{
		User: env["CLIENT_USERNAME"],
		Auth: []ssh.AuthMethod{
			ssh.Password(env["CLIENT_PASSWORD"]),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", env["SERVER_ADDRESS"], conf)
	if err != nil {
		lg.Error(err.Error())
		return
	}
	defer conn.Close()

	c := &client{
		config:     conf,
		connection: conn,
		logger:     lg,
	}
	c.start()
}
