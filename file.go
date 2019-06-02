package main

import "time"

type File struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Data       string    `json:"data"`
	Public     bool      `json:"public"`
	UploadDate time.Time `json:"upload_date"`
}

type Files []File
