package entity

import "time"

type Post struct {
    ID    int `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    Date time.Time `json:"date"`
}