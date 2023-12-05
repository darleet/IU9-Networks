package main

import (
	_ "embed"
	"html/template"
	"net/http"
	"newsfeed/src/service"

	log "github.com/mgutz/logxi/v1"
)

//go:embed templates/index.html
var indexHTML string
var tIndex = template.Must(template.New("index.html").Parse(indexHTML))

func serveClient(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	log.Info("got request", "Method", r.Method, "Path", path)
	if path != "/" && path != "/index.html" {
		log.Error("invalid path", "Path", path)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := service.DownloadNews()
	if data == nil {
		log.Error("Empty data", "error", data)
		return
	}
	if err := tIndex.Execute(w, data); err != nil {
		log.Error("HTML creation failed", "error", err)
		return
	}
	log.Info("response sent to client successfully")
}

func main() {
	http.HandleFunc("/", serveClient)
	log.Info("starting listener")
	log.Error("listener failed", "error", http.ListenAndServe("127.0.0.1:9000", nil))
}
