package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/messages"
	"webapp/src/models"
	"webapp/src/requisicoes"
	"webapp/src/utils"
)

// CarregarTelaDeLogin -> render the login screen
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Login screen"))
	cookie, _ := cookies.Ler(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

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

	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []models.Publicacao
		UsuarioId   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioId:   usuarioId,
	})
}

// CarregarPaginaEditarPost -> Load page post edit
func CarregarPaginaEditarPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.ApiUrl, postId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	var post models.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&post); erro != nil {
		messages.JSON(w, http.StatusUnprocessableEntity, messages.ErroApi{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "post-edit.html", post)
}

// CarregarPaginaUsuarios -> Search users
func CarregarPaginaUsuarios(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.ApiUrl, nameOrNick)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	var usuarios []models.Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		messages.JSON(w, http.StatusUnprocessableEntity, messages.ErroApi{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilUsuario -> Load user perfil
func CarregarPerfilUsuario(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parameters["usuarioId"], 10, 64)
	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioId == usuarioLogadoId {
		http.Redirect(w, r, "/perfil", 302)
		return
	}

	usuario, erro := models.BuscarUsuarioCompleto(usuarioId, r)
	// fmt.Println(usuario, erro)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         models.Usuario
		UsuarioLogadoId uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoId: usuarioLogadoId,
	})

}

// CarregarPerfilUsuarioLogado -> Carregar perfil usuário logado
func CarregarPerfilUsuarioLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	usuario, erro := models.BuscarUsuarioCompleto(usuarioId, r)
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", usuario)
}
