package main

import (
	"gopkg.in/ini.v1"
	"lab5/client"
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
		log.Fatal("no address provided in .ini")
	}
	client.Run(addr)
}
