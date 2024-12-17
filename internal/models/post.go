package models // package declaration, that this file bleongs to models package

import "gorm.io/gorm"

type Post struct {
	gorm.Model // optional, this will add ID, CreatedAt, UpdatedAt, DeletedAt fields to the model
	ID string `json:"id"` // this is metadata for Go JSON encoder/decoder to convert
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
}