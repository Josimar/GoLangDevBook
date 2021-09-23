package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/messages"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// CriarPublicacao - create a post
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		messages.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.UserId = usuarioId

	if erro = publicacao.Preparar(); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusCreated, publicacao)
}

// BuscarPublicacoes - list posts
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	post, erro := repositorio.Buscar(usuarioId)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusOK, post)
}

// BuscarPublicacao - list a post
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	postId, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	post, erro := repositorio.BuscarPorId(postId)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusOK, post)
}

// AtualizarPublicacao - update a post
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.UserId != usuarioId {
		messages.Erro(w, http.StatusForbidden, errors.New("post can't be to alter for another user"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		messages.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao models.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacao.Preparar(); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicacaoId, publicacao); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// ExcluirPublicacao - delete a post
func ExcluirPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	publicacaoSalvaNoBanco, erro := repositorio.BuscarPorId(publicacaoId)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoSalvaNoBanco.UserId != usuarioId {
		messages.Erro(w, http.StatusForbidden, errors.New("post can't be to alter for another user"))
		return
	}

	if erro = repositorio.Deletar(publicacaoId); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// BuscarPublicacoesPorUsuario - list posts of users
func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	postId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	post, erro := repositorio.BuscarPorUsuario(postId)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusOK, post)
}

// CurtirPublicacao - like a post
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	if erro = repositorio.Curtir(publicacaoId); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// DisCurtirPublicacao - like a post
func DisCurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["postId"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewPostRepository(db)
	if erro = repositorio.DisCurtir(publicacaoId); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}
