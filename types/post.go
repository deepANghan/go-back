package types

import "time"

type CreatePost struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type Post struct {
	Id        int16     `json:id`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Timestamp time.Time `json:"timestamp"`
}
