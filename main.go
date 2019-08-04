package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Println("Goupfile server is starting...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
