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
func (repository usuarios) Create(usuario models.Usuario) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO users (name, email, password, created_at) VALUES(?, ?, ?, ?)")
	if erro != nil{
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Name, usuario.Email, usuario.Password, usuario.CreatedAt)
	if erro != nil{
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil{
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
