package main

import (
	"log"
	"websockets/client"

	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}
	addr, ok := env["CONNECTION_ADDRESS"]
	if !ok {
		log.Fatal("no connection address provided in .env")
	}
	client.Run(addr)
}
