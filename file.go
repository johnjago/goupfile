package main

import "time"

type File struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	MediaType  string    `json:"media_type"`
	UploadDate time.Time `json:"upload_date"`
}

type Files []File
