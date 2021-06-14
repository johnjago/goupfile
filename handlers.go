package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// handleUpload implements the POST request for uploading a file.
func handleUpload(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	io.Copy(&Buf, file)

	id := generateID(6)

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		log.Fatal(err)
	}

	err = ioutil.WriteFile("uploads/"+id, Buf.Bytes(), os.ModePerm)
	if err != nil {
		errMsg := "Unable to save file."
		log.Println(errMsg, err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// I reset the buffer in case I want to use it again.
	Buf.Reset()

	fileData := File{
		ID:         id,
		Name:       header.Filename,
		Size:       header.Size,
		MediaType:  header.Header.Get("Content-Type"),
		UploadDate: time.Now(),
	}
	// TODO: Put this back after switching to SQLite
	// f := DBCreateFile(fileData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(fileData)
}
