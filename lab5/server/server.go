package server

import (
	"log"
	"net/http"
	"websockets/server/service"
)

func Run(address string) {
	http.HandleFunc("/", service.HandleSocket)
	log.Fatal(http.ListenAndServe(address, nil))
}
