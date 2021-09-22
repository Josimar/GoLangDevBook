package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ExcluirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}/follower",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}/followers",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaSeguidores,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{id}/follow",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscaSeguindo,
		RequerAutenticacao: true,
	},
}
