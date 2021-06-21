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

const schema = `create table if not exists Files (
	ID varchar(10) not null primary key,
	Name varchar(255) not null,
	Size integer not null,
	MediaType varchar(255) not null,
	UploadDate timestamp
);`

func main() {
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/upload", Logger(handleUpload, "handleUpload"))
	http.HandleFunc("/download", Logger(handleDownload, "handleDownload"))

	initDB(driver, dataSource, schema)
	log.Printf("Goupfile starting on %s%s\n", host, port)
	log.Fatal(http.ListenAndServe(port, nil))
}
