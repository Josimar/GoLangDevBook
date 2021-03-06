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

// BuscarPorId -> traz um usuário pelo ID
func (repository usuarios) BuscarPorId(Id uint64) (models.Usuario, error) {
	linhas, erro := repository.db.Query("SELECT id, name, email, created_at FROM users WHERE id = ?", Id)
	if erro != nil {
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Name,
			&usuario.Email,
			&usuario.CreatedAt,
		); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar -> atualiza o cadastro do usuário
func (repository usuarios) Atualizar(Id uint64, usuario models.Usuario) error {
	statement, erro := repository.db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuario.Name, usuario.Email, Id); erro != nil {
		return erro
	}

	return nil
}

// Deletar -> exclui o usuário da base de dados
func (repository usuarios) Deletar(Id uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Id); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail -> Search a user by email
func (repository usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Password); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Seguir -> check one user follow another user
func (repository usuarios) Seguir(usuarioId, seguidorId uint64) error {
	statement, erro := repository.db.Prepare("INSERT INTO followers(user_id, follower_id) VALUES(?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

// DeSeguir -> uncheck one user follow another user
func (repository usuarios) DeSeguir(usuarioId, seguidorId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

// BuscarSeguidores -> busca os seguidores de um usuário
func (repository usuarios) BuscarSeguidores(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := repository.db.Query("SELECT U.id, U.name, U.email FROM users U INNER JOIN followers F ON U.id = F.follower_id WHERE F.user_id = ?", usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Name, &usuario.Email); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSeguindo -> busca usuários que um usuário está seguindo
func (repository usuarios) BuscarSeguindo(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := repository.db.Query("SELECT U.id, U.name, U.email FROM users U INNER JOIN followers F ON U.id = F.user_id WHERE F.follower_id = ?", usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Name, &usuario.Email); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarSenha -> busca senha do usuário logado
func (repository usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linhas, erro := repository.db.Query("SELECT password FROM users WHERE id = ?", usuarioId)
	if erro != nil {
		return "", erro
	}
	defer linhas.Close()

	var usuario models.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(&usuario.Password); erro != nil {
			return "", erro
		}
	}

	return usuario.Password, nil
}

// AtualizarSenha -> update the password of the user
func (repository usuarios) AtualizarSenha(usuarioId uint64, senha string) error {
	statement, erro := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, usuarioId); erro != nil {
		return erro
	}

	return nil
}
