package main

import (
	"gopkg.in/ini.v1"
	"lab5/server"
	"log"
)

func main() {
	cfg, err := ini.Load(".ini")
	if err != nil {
		log.Fatal(err)
	}
	data := cfg.Section("connection")
	addr := data.Key("address").String()
	if addr == "" {
		addr = "localhost:8080"
	}
	server.Run(addr)
}
