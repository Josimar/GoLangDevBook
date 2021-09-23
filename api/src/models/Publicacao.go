package models

import (
	"errors"
	"strings"
	"time"
)

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

// Preparar -> control fields
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()
	return nil
}

// validar -> check empty fields
func (publicacao *Publicacao) validar() error {
	if publicacao.Title == "" {
		return errors.New("the title is required")
	}

	if publicacao.Content == "" {
		return errors.New("the content is required")
	}

	if publicacao.Description == "" {
		return errors.New("the description is required")
	}

	return nil
}

// formatar -> format content
func (publicacao *Publicacao) formatar() {
	publicacao.Title = strings.TrimSpace(publicacao.Title)
	publicacao.Content = strings.TrimSpace(publicacao.Content)
	publicacao.Description = strings.TrimSpace(publicacao.Description)
}
