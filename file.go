package main

import "time"

type File struct {
	ID         string    `json:"id"`
	Group      string    `json:"group"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	MediaType  string    `json:"media_type"`
	UploadDate time.Time `json:"upload_date"`
	URL        string    `json:"url"`
}

type FileGroup struct {
	ID    string `json:"id"`
	Files []File `json:"files"`
}
