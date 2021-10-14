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

// CurtirPost -> check post with favorite
func CurtirPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.ApiUrl, postId)
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

// DesCurtirPost -> uncheck post with favorite
func DesCurtirPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.ApiUrl, postId)
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
