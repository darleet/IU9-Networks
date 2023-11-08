package main

import (
	"log/slog"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh/terminal"
)

func handler(s ssh.Session) {
	term := terminal.NewTerminal(s, "> ")
	for {
		line, err := term.ReadLine()
		if err != nil {
			break
		}
		slog.Info(line)
	}
	slog.Info("terminal closed")
}

func main() {
	env, err := godotenv.Read()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	ssh.Handle(handler)
	if err := ssh.ListenAndServe(env["SERVER_HOST_ADDRESS"], nil,
		ssh.HostKeyFile(env["SERVER_HOST_KEY"])); err != nil {
		slog.Error(err.Error())
		return
	}
}
