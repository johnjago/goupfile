package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Goupfile! Send an HTTP POST request to upload a file.")
}

func FileShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// TODO: Take ID in the format ad3D9a, look up in DB to get file metadata
	// mediaType := "image/png"
	// w.Header().Set("Content-Type", mediaType + "; charset=utf-8")
	// w.Header().Set("Content-Disposition", "attachment; filename=" + vars["id"])
	http.ServeFile(w, r, "uploads/"+vars["id"])
}

func FileCreate(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	io.Copy(&Buf, file)
	err = ioutil.WriteFile("uploads/"+header.Filename, Buf.Bytes(), 0644)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// I reset the buffer in case I want to use it again.
	// Reduces memory allocations in more intense projects.
	Buf.Reset()

	fileData := File{
		Name:       header.Filename,
		MediaType:  header.Header.Get("Content-Type"),
		UploadDate: time.Now(),
	}
	f := DBCreateFile(fileData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(f); err != nil {
		http.Error(w, "Unable to get file JSON from struct", http.StatusInternalServerError)
		return
	}
}
