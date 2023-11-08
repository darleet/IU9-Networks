package main

import "github.com/gliderlabs/ssh"

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
