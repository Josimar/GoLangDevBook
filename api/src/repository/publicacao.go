package repository

import (
	"api/src/models"
	"database/sql"
	"time"
)

// Publicacao -> represent a post
type publicacao struct {
	db *sql.DB
}

// NewPostRepository -> create a repository of post
func NewPostRepository(db *sql.DB) *publicacao {
	return &publicacao{db}
}

// Criar -> create a post in the database
func (repository publicacao) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO posts(title, content, description, user_id) VALUES(?, ?, ?, ?) ")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Title, publicacao.Content, publicacao.Description, publicacao.UserId)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// BuscarPorId -> search for id
func (repository publicacao) BuscarPorId(postId uint64) (models.Publicacao, error) {
	linha, erro := repository.db.Query("SELECT p.id, p.title, p.content, p.description, p.user_id, IFNULL(p.likes, 0) as likes, p.created_at, u.name FROM posts p INNER JOIN users u ON u.id = p.user_id WHERE p.id = ?", postId)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao models.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Title,
			&publicacao.Content,
			&publicacao.Description,
			&publicacao.UserId,
			&publicacao.Likes,
			&publicacao.CreatedAt,
			&publicacao.Username,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar -> search for id
func (repository publicacao) Buscar(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repository.db.Query("SELECT p.id, p.title, p.content, p.description, p.user_id, IFNULL(p.likes, 0) as likes, p.created_at, u.name FROM posts p INNER JOIN users u ON u.id = p.user_id INNER JOIN followers f ON p.user_id = f.user_id WHERE u.id = ? OR f.follower_id = ? ORDER BY p.id DESC", usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Title,
			&publicacao.Content,
			&publicacao.Description,
			&publicacao.UserId,
			&publicacao.Likes,
			&publicacao.CreatedAt,
			&publicacao.Username,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar -> update a post
func (repository publicacao) Atualizar(publicacaoId uint64, publicacao models.Publicacao) error {
	statement, erro := repository.db.Prepare("UPDATE posts SET title = ?, content = ?, description = ?, updated_at = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Title, publicacao.Content, publicacao.Description, time.Now(), publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// Deletar -> update a post
func (repository publicacao) Deletar(publicacaoId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUsuario -> search posts for user
func (repository publicacao) BuscarPorUsuario(usuarioId uint64) ([]models.Publicacao, error) {
	linhas, erro := repository.db.Query("SELECT p.id, p.title, p.content, p.description, p.user_id, IFNULL(p.likes, 0) as likes, p.created_at, u.name FROM posts p INNER JOIN users u ON u.id = p.user_id INNER JOIN followers f ON p.user_id = f.user_id WHERE u.id = ? ORDER BY p.id DESC", usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []models.Publicacao

	for linhas.Next() {
		var publicacao models.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Title,
			&publicacao.Content,
			&publicacao.Description,
			&publicacao.UserId,
			&publicacao.Likes,
			&publicacao.CreatedAt,
			&publicacao.Username,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir -> update a post
func (repository publicacao) Curtir(publicacaoId uint64) error {
	// ToDo: like tem de ser em outra tabela para o mesmo user não ficar marcando como like
	statement, erro := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}

// DisCurtir -> update a post
func (repository publicacao) DisCurtir(publicacaoId uint64) error {
	// ToDo: like tem de ser em outra tabela para o mesmo user não ficar desmarcando como like
	statement, erro := repository.db.Prepare("UPDATE posts SET likes = likes - 1 WHERE id = ? AND likes > 0")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoId); erro != nil {
		return erro
	}

	return nil
}
