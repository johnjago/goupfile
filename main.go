package main

import (
	"log"
	"net/http"
)

const (
	port       = ":8090"
	staticDir  = "./public"
	driver     = "sqlite3"
	dataSource = "sqlite_db"
	// This is only for generating the QR code from a URL since when it's
	// running behind nginx, for example, it doesn't know what the outside
	// facing hostname is unless we tell it. A real setup might look like:
	// this program running on localhost:8090, but nginx accepts requests at
	// https://file.com, in which case the below value would be https://file.com
	publicHost = "https://goupfile.johnjago.com"
)

const schema = `
create table if not exists files (
	id varchar(10) not null primary key,
	group_id varchar(10) not null,
	name varchar(255) not null,
	size integer not null,
	media_type varchar(255) not null,
	upload_date timestamp,
	url varchar(255) not null
);
`

func main() {
	http.Handle("/", http.FileServer(http.Dir(staticDir)))
	http.HandleFunc("/upload", Logger(handleUpload))
	http.HandleFunc("/d/", Logger(handleDownload))
	http.HandleFunc("/v/", Logger(handleView))

	if err := initDB(driver, dataSource, schema); err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	log.Printf("Goupfile starting on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
