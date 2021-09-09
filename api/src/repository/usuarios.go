package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Name, usuario.Email, usuario.Password, usuario.CreatedAt)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Buscar -> traz todos usuários pelo nome ou email
func (repository usuarios) Buscar(nomeOuEmail string) ([]models.Usuario, error) {
	nomeOuEmail = fmt.Sprintf("%%%s%%", nomeOuEmail)

	linhas, erro := repository.db.Query("SELECT id, name, email, created_at FROM users WHERE name LIKE ? OR email LIKE ?", nomeOuEmail, nomeOuEmail)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Email,
			&usuario.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}