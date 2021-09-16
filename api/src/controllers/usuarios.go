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
	"strings"
)

// CriarUsuario -> chamada de criação do usuário pelo controller
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		messages.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	usuarioId, erro := repositorio.Create(usuario)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario.ID = usuarioId
	messages.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios -> busca todos usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuEmail := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	usuarios, erro := repositorio.Buscar(nomeOuEmail)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario -> busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["id"], 10, 64)
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

	repositorio := repository.NewUserRepository(db)
	usuario, erro := repositorio.BuscarPorId(usuarioID)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario -> atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// fmt.Println(usuarioIDNoToken)
	if usuarioId != usuarioIDNoToken {
		messages.Erro(w, http.StatusForbidden, errors.New("action not allowed"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		messages.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("atualizar"); erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	erro = repositorio.Atualizar(usuarioId, usuario)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// ExcluirUsuario -> exclui um usuário
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// fmt.Println(usuarioIDNoToken)
	if usuarioId != usuarioIDNoToken {
		messages.Erro(w, http.StatusForbidden, errors.New("action not allowed"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	if erro = repositorio.Deletar(usuarioId); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// SeguirUsuario -> follower a user
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorId == usuarioId {
		messages.Erro(w, http.StatusForbidden, errors.New("you can not follow yourself"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	if erro = repositorio.Seguir(usuarioId, seguidorId); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}

// DesSeguirUsuario -> Unfollow a user
func DesSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorId, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["id"], 10, 64)
	if erro != nil {
		messages.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidorId == usuarioId {
		messages.Erro(w, http.StatusForbidden, errors.New("you can not follow yourself"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	if erro = repositorio.DeSeguir(usuarioId, seguidorId); erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	messages.JSON(w, http.StatusNoContent, nil)
}
