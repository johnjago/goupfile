package main

import (
	"log"
	"net/http"
)

const (
	scheme     = "http"
	host       = "localhost"
	port       = ":8090"
	staticDir  = "./public"
	driver     = "sqlite3"
	dataSource = "sqlite_db"
)

// TODO: Remove drop table after schema is finalized
const schema = `
drop table Files;

create table Files (
	ID varchar(10) not null primary key,
	Name varchar(255) not null,
	Size integer not null,
	MediaType varchar(255) not null,
	UploadDate timestamp,
	URL varchar(255) not null
);`

func main() {
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/upload", Logger(handleUpload))
	http.HandleFunc("/d/", Logger(handleDownload))
	http.HandleFunc("/v/", Logger(handleView))

	initDB(driver, dataSource, schema)
	log.Printf("Goupfile starting on %s%s\n", host, port)
	log.Fatal(http.ListenAndServe(port, nil))
}
