package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/messages"
	"api/src/models"
	"api/src/repository"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	token, erro := authentication.CreateToken(usuarioSalvoBanco.ID)
	if erro != nil {
		messages.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	fmt.Println(token)
	// w.Write([]byte(token))

	usuarioId := strconv.FormatUint(usuarioSalvoBanco.ID, 10)

	messages.JSON(w, http.StatusOK, models.DadosAutenticacao{Id: usuarioId, Token: token})
}
