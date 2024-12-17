package models // package declaration, that this file bleongs to models package

type Post struct {
	ID string `json:"id"` // this is metadata for Go JSON encoder/decoder to convert
	Title string `json:"title"`
	Body string `json:"body"`
	Author string `json:"author"`
}