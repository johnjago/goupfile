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

// todo: remove drop table after schema is finalized
const schema = `
drop table if exists files;

create table files (
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

	log.Printf("Goupfile starting on %s%s\n", host, port)
	log.Fatal(http.ListenAndServe(port, nil))
}
