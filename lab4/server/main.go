package main

import (
	"io"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh/terminal"
)

// handleSSH is the main handler for ssh
func handleSSH(s ssh.Session) {
	term := terminal.NewTerminal(s, "> ")
	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			slog.Info("ssh client exited session.")
			return
		} else if err != nil {
			slog.Error(err.Error())
			continue
		}
		fields := strings.Fields(line)
		cmd := exec.Command(fields[0], fields[1:]...)
		cmd.Stdout = s
		cmd.Stderr = s
		if err := cmd.Run(); err != nil {
			slog.Error(err.Error())
		}
	}
}

func main() {
	env, err := godotenv.Read()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	a := auth{
		users: map[string]string{
			"darleet": "test1234",
		},
	}

	server := ssh.Server{
		Addr:            env["SERVER_HOST_ADDRESS"],
		Handler:         handleSSH,
		PasswordHandler: a.handlePassword,
	}
	server.SetOption(ssh.HostKeyFile(env["SERVER_HOST_KEY"]))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		return
	}
}
