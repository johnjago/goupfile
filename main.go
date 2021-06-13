package main

import (
	"log"
	"net/http"
)

const (
	staticDir = "./public"
	port      = ":8090"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/upload", handleUpload)

	log.Printf("Goupfile starting on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
