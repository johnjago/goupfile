package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(driver, dataSource, schema string) error {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return err
	}
	if _, err = db.Exec(schema); err != nil {
		return err
	}
	return nil
}

// saveFile saves file metadata into the database. If successful, it returns
// the same file struct.
func saveFile(f File) File {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}
	insert := "insert into Files (ID, Name, Size, MediaType) values (?, ?, ?, ?)"
	_, err = db.Exec(insert, f.ID, f.Name, f.Size, f.MediaType)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// getFile obtains information about a specific file from the database.
// It does not obtain the file itself, as that is stored on the file system.
func getFile(id string) File {
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
