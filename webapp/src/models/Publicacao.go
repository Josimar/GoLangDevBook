package models

import "time"

// Publicacao -> table of posts
type Publicacao struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Content     string    `json:"content,omitempty"`
	UserId      uint64    `json:"user_id,omitempty"`
	Username    string    `json:"user_name,omitempty"`
	Likes       uint64    `json:"likes,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
