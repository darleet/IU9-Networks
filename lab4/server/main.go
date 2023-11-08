package main

import (
	"io"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh/terminal"
)

// auth represents authentication service
type auth struct {
	// users stores [login]password pairs
	users map[string]string
}

// handlePassword handles password authentication for ssh
func (a *auth) handlePassword(ctx ssh.Context, password string) bool {
	pass, exists := a.users[ctx.User()]
	return exists && pass == password
}

// handleSFTP handles SFTP connection (за это обещали доп балл!!)
func handleSFTP(sess ssh.Session) {
	server, err := sftp.NewServer(sess)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if err := server.Serve(); err == io.EOF {
		server.Close()
		slog.Info("sftp client exited session.")
	} else if err != nil {
		slog.Error(err.Error())
		return
	}
}

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
		SubsystemHandlers: map[string]ssh.SubsystemHandler{
			"sftp": handleSFTP,
		},
	}
	server.SetOption(ssh.HostKeyFile(env["SERVER_HOST_KEY"]))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
		return
	}
}
