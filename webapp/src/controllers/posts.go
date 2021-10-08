package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/messages"
	"webapp/src/requisicoes"
)

// CriarPosts -> chama API para criar um post
func CriarPosts(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"title":       r.FormValue("title"),
		"content":     r.FormValue("content"),
		"description": r.FormValue("description"),
	})
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.ApiUrl)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
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
