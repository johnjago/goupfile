package main

import "fmt"

var currentId int

var files Files

func init() {
	DBCreateFile(File{Name: "one.txt"})
	DBCreateFile(File{Name: "two.txt"})
}

func DBFindFile(id int) File {
	for _, f := range files {
		if f.Id == id {
			return f
		}
	}
	return File{}
}

func DBCreateFile(f File) File {
	currentId++
	f.Id = currentId
	files = append(files, f)
	return f
}

func DBDeleteFile(id int) error {
	for i, f := range files {
		if f.Id == id {
			files = append(files[:i], files[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find file with id of %d to delete", id)
}
