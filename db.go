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
	insert := "insert into files (id, group_id, name, size, media_type, url) values (?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(insert, f.ID, f.Group, f.Name, f.Size, f.MediaType, f.URL)
	if err != nil {
		log.Fatal("Failed to insert file into files table: ", err)
	}
	return f
}

// getFile obtains information about a specific file from the database.
// It does not obtain the file itself, as that is stored on the file system.
func getFile(id string) File {
	// todo: only call sql.Open since it has its own connection pool, instead of
	// here and in the other function
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, group_id, name, size, media_type, url from files where id = ?", id)
	if err != nil {
		log.Printf("Failed to get file %s: %s", id, err)
	}
	defer rows.Close()

	f := File{}

	for rows.Next() {
		if err := rows.Scan(&f.ID, &f.Group, &f.Name, &f.Size, &f.MediaType, &f.URL); err != nil {
			log.Fatal(err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return f
}

// getFileGroup returns a struct representing a group of files saved under a
// single URL.
func getFileGroup(id string) FileGroup {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select id, group_id, name, size, media_type, url from files where group_id = ?", id)
	if err != nil {
		log.Printf("Failed to get files in group %s: %s", id, err)
	}
	defer rows.Close()

	fileGroup := FileGroup{ID: id}

	for rows.Next() {
		f := File{}
		if err := rows.Scan(&f.ID, &f.Group, &f.Name, &f.Size, &f.MediaType, &f.URL); err != nil {
			log.Fatal(err)
		}
		fileGroup.Files = append(fileGroup.Files, f)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return fileGroup
}
