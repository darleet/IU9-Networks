package main

import (
	"io"
	"lab4/server/service"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"github.com/pkg/sftp"
	"golang.org/x/term"
)

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
	t := term.NewTerminal(s, "> ")
	for {
		line, err := t.ReadLine()
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

	a := service.Auth{
		Users: map[string]string{
			"darleet": "test1234",
		},
	}

	server := ssh.Server{
		Addr:            env["SERVER_HOST_ADDRESS"],
		Handler:         handleSSH,
		PasswordHandler: a.HandlePassword,
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
