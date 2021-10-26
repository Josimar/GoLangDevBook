package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/messages"
	"webapp/src/requisicoes"
)

// CriarUsuario -> chamar API para criar usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//name := r.FormValue("name")
	// fmt.Println(name)

	usuario, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	// fmt.Println(usuario)
	// fmt.Println(bytes.NewBuffer(usuario))

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// fmt.Println(response.Body)

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	messages.JSON(w, response.StatusCode, nil)
}

// PararDeSeguirUsuario -> Para de seguir o usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/unfollow", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	messages.JSON(w, response.StatusCode, nil)
}

// SeguirUsuario -> Seguir usuário
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/follower", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	messages.JSON(w, response.StatusCode, nil)
}
