package main

import "fmt"

var files Files

func init() {
	DBCreateFile(File{Name: "one.txt"})
}

func DBFindFile(id string) File {
	for _, f := range files {
		if f.ID == id {
			return f
		}
	}
	return File{}
}

func DBCreateFile(f File) File {
	files = append(files, f)
	return f
}

func DBDeleteFile(id string) error {
	for i, f := range files {
		if f.ID == id {
			files = append(files[:i], files[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find file with id of %s to delete", id)
}
