package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/messages"
	"webapp/src/models"
)

// FazerLogin -> faz o login
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// token, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(response.StatusCode, response.Body)
	// fmt.Println(response.StatusCode, string(token))

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao models.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		messages.JSON(w, http.StatusUnprocessableEntity, messages.ErroApi{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.Id, dadosAutenticacao.Token); erro != nil {
		messages.JSON(w, http.StatusUnprocessableEntity, messages.ErroApi{Erro: erro.Error()})
		return
	}

	messages.JSON(w, http.StatusOK, nil)
}
