package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func DBCreateFile(f File) File {
	// TODO
	return f
}

// DBGetFile obtains information about a specific file from the database.
// It does not obtain the file itself, as that is stored on the file system.
func DBGetFile(id string) File {
	db, err := sql.Open("mysql", "goupfile:password@/goupfile")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select ID, Name, MediaType from Files")

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
