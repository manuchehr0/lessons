package models

import "time"

//type User struct {
//	Id
//	Name
//}

type Note struct {
	Id        int64      `json:"id"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
