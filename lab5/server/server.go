package server

import (
	"log"
	"net/http"
)

func Run(address string) {
	http.HandleFunc("/", handleSocket)
	log.Fatal(http.ListenAndServe(address, nil))
}
