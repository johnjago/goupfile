package main

import "time"

type File struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	MediaType  string    `json:"media_type"`
	UploadDate time.Time `json:"upload_date"`
}

type Files []File
