package controllers

import (
	"api/src/banco"
	"api/src/messages"
	"api/src/models"
	"api/src/repository"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login -> authenticate the user
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := banco.Conectar()
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NewUserRepository(db)
	usuarioSalvoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuarioSalvoBanco.Password, usuario.Password); erro != nil {
		messages.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	w.Write([]byte("LogIn with successfully"))
}
