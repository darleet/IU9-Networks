package main

import (
	"log"
	"websockets/server"

	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	addr, ok := env["CONNECTION_ADDRESS"]
	if !ok {
		addr = "localhost:8080"
	}
	server.Run(addr)
}
