package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
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

	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		log.Println(err)
		log.Println("Making a new uploads directory")
		if err := os.Mkdir("uploads", os.ModePerm); err != nil {
			message := "Could not create uploads directory"
			log.Println(message, err)
			http.Error(w, message, http.StatusInternalServerError)
		}
	}

	id := generateID(6)
	err = ioutil.WriteFile("uploads/"+id, Buf.Bytes(), os.ModePerm)
	if err != nil {
		message := "Unable to save file"
		log.Println(message, err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	Buf.Reset()

	fileData := File{
		ID:         id,
		Name:       header.Filename,
		Size:       header.Size,
		MediaType:  header.Header.Get("Content-Type"),
		UploadDate: time.Now(),
		URL:        downloadURL(r, id),
	}
	log.Println("Uploaded file:", fileData)
	f := saveFile(fileData)

	http.Redirect(w, r, viewURL(r, f.ID), http.StatusSeeOther)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, present := query["id"]
	if !present || len(id) == 0 {
		log.Println("File ID not present")
	}
	f := getFile(id[0])
	w.Header().Set("Content-Type", f.MediaType+"; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename="+f.Name)
	http.ServeFile(w, r, "uploads/"+f.ID)
}

func handleView(w http.ResponseWriter, r *http.Request) {
	file := getFile(r.URL.Path[len("/f/"):])
	t, _ := template.ParseFiles("templates/view.html")
	t.Execute(w, file)
}

func viewURL(r *http.Request, id string) string {
	return scheme + "://" + host + port + "/f/" + id
}

func downloadURL(r *http.Request, id string) string {
	return scheme + "://" + host + port + "/download?id=" + id
}
