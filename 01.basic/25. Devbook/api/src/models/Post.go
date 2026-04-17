package models

import (
	"errors"
	"strings"
	"time"
)

// Post represents a post in the system
type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"author_id,omitempty"`
	AuthorNick string    `json:"author_nick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

// PreparePost will format and validate the post data
func (post *Post) PreparePost() error {
	if err := post.validate(); err != nil {
		return err
	}
	post.formatter()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title is not be blank")
	}
	if post.Content == "" {
		return errors.New("content is not be blank")
	}
	return nil
}

func (post *Post) formatter() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
