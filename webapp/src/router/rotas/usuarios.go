package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaCadastroUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/buscar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuario/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuario/{usuarioId}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuario/{usuarioId}/follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
}
