package repository

import (
	"api/src/models"
	"database/sql"
)

// struct de usuário - local
type usuarios struct {
	db *sql.DB
}

// NewUserRepository -> Criar um repositório de usuário
func NewUserRepository(db *sql.DB) *usuarios {
	return &usuarios{db}
}

// Create -> insere um usuário na base de dados
func (u usuarios) Create(usuario models.Usuario) (uint64, error) {
	return 0, nil
}
