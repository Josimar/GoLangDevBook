package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario -> representa uma pessoa
type Usuario struct {
	ID        uint64       `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	CreatedAt time.Time    `json:"created_at"`
	Followers []Usuario    `json:"followers"`
	Follows   []Usuario    `json:"follows"`
	Posts     []Publicacao `json:"posts"`
}

// BuscarUsuarioCompleto -> busca 4 requisições a API
func BuscarUsuarioCompleto(usuarioId uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalFollowers := make(chan []Usuario)
	canalFollows := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioId, r)
	go BuscarFollowers(canalFollowers, usuarioId, r)
	go BuscarFollows(canalFollows, usuarioId, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioId, r)

	var (
		usuario     Usuario
		followers   []Usuario
		follows     []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("Erro ao buscar o usuário")
			}
			usuario = usuarioCarregado

		case followersCarregado := <-canalFollowers:
			if followersCarregado == nil {
				return Usuario{}, errors.New("Erro ao buscar os followers")
			}
			followers = followersCarregado

		case followsCarregado := <-canalFollows:
			if followsCarregado == nil {
				return Usuario{}, errors.New("Erro ao buscar os follows")
			}
			follows = followsCarregado

		case postCarregado := <-canalPublicacoes:
			if postCarregado == nil {
				return Usuario{}, errors.New("Erro ao buscar os follows")
			}
			publicacoes = postCarregado
		}
	}

	usuario.Followers = followers
	usuario.Follows = follows
	usuario.Posts = publicacoes

	return usuario, nil
}

// BuscarDadosDoUsuario -> Search user
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

// BuscarFollowers -> search followers (seguidores)
func BuscarFollowers(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/followers", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var followers []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&followers); erro != nil {
		canal <- nil
		return
	}

	if followers == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- followers
}

// BuscarFollows -> search follows - seguidos
func BuscarFollows(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/follow", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var follows []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&follows); erro != nil {
		canal <- nil
		return
	}

	if follows == nil {
		canal <- make([]Usuario, 0)
		return
	}

	canal <- follows
}

// BuscarPublicacoes -> search post
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/posts", config.ApiUrl, usuarioId)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var posts []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&posts); erro != nil {
		canal <- nil
		return
	}

	if posts == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- posts
}
