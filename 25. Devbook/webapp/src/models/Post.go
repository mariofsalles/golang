package models

import "time"

// Post represents the post made by the user
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      int       `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
