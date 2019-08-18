package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/goupfile/core/templates"
	"github.com/gorilla/mux"
)

// Index is the default response for a GET request without an ID.
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, templates.Index)
}

// FileShow implements the GET request for downloading a file given an ID.
func FileShow(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)
	f := DBGetFile(urlVars["id"])
	w.Header().Set("Content-Type", f.MediaType+"; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename="+f.Name)
	http.ServeFile(w, r, "uploads/"+f.ID)
}

// FileCreate implements the POST request for uploading a file.
func FileCreate(w http.ResponseWriter, r *http.Request) {
	var Buf bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	io.Copy(&Buf, file)

	id := generateID(6)
	err = ioutil.WriteFile("uploads/"+id, Buf.Bytes(), 0644)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	// I reset the buffer in case I want to use it again.
	// Reduces memory allocations in more intense projects.
	Buf.Reset()

	fileData := File{
		ID:         id,
		Name:       header.Filename,
		MediaType:  header.Header.Get("Content-Type"),
		UploadDate: time.Now(),
	}
	f := DBCreateFile(fileData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	// TODO: Check .env for prod/local
	fmt.Fprintf(w, "\n\thttps://goupfile.com/%s\n\n", f.ID)
}
