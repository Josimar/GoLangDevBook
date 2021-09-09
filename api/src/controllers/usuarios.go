package controllers

import (
	"api/src/banco"
	"api/src/messages"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	if erro = usuario.Preparar(); erro != nil {
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
	w.Write([]byte("Controllers -> BuscarUsuario -> Buscar um usuário"))
}

// AtualizarUsuario -> atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> AtualizarUsuario -> Atualizar um usuário"))
}

// ExcluirUsuario -> exclui um usuário
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> ExcluirUsuario -> Excluir um usuário"))
}
