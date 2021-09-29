package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/messages"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/utils"
)

// CarregarTelaDeLogin -> render the login screen
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Login screen"))
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaCadastroUsuario -> render register page
func CarregarPaginaCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal -> render main page
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.ApiUrl)

	// response, erro := http.Get(url)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode, erro)

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacoes []models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		messages.JSON(w, http.StatusUnprocessableEntity, messages.ErroApi{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []models.Publicacao
		OutroCampo  string
	}{
		Publicacoes: publicacoes,
		OutroCampo:  "Outro campo",
	})
}
