package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	driver     string = "mysql"
	dataSource string = "goupfile:password@/goupfile"
)

// DBCreateFile saves file metadata into the database. If successful, it returns
// the same file struct.
func DBCreateFile(f File) File {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	insert := "insert into Files (ID, Name, MediaType) values (?, ?, ?)"
	_, err = db.Exec(insert, f.ID, f.Name, f.MediaType)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// DBGetFile obtains information about a specific file from the database.
// It does not obtain the file itself, as that is stored on the file system.
func DBGetFile(id string) File {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select ID, Name, MediaType from Files where ID = ?", id)
	f := File{}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&f.ID, &f.Name, &f.MediaType); err != nil {
			log.Fatal(err)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return f
}
