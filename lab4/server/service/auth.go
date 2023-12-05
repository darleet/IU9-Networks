package service

import "github.com/gliderlabs/ssh"

// Auth represents authentication service
type Auth struct {
	// Users stores [login]password pairs
	Users map[string]string
}

// HandlePassword handles password authentication for ssh
func (a *Auth) HandlePassword(ctx ssh.Context, password string) bool {
	pass, exists := a.Users[ctx.User()]
	return exists && pass == password
}
