package main

import (
	"fmt"
	"log"
	"net/http"
)

func ContactRouterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contacts/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./static/contacts.html")
}

func InfoRouterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/info/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "./static/info.html")
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./static/index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "400 Bad request", http.StatusBadRequest)
			return
		}
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", v)
		}
		fmt.Fprintf(w, "Data sent!")
	}
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/info/", InfoRouterHandler)
	http.HandleFunc("/contacts/", ContactRouterHandler)
	if err := http.ListenAndServe("localhost:9000", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
