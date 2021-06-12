package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	staticDir = "./public"
	port      = ":8090"
)

type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/api/v1/upload", handleUpload)

	log.Printf("Goupfile starting on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	f := File{
		Name: handler.Filename,
		Size: handler.Size,
	}

	json.NewEncoder(w).Encode(f)
}
