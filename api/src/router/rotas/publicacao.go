package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacao = []Rota{
	{
		Uri:                "/posts",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ExcluirPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioId}/posts",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}/like",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}/dislike",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DisCurtirPublicacao,
		RequerAutenticacao: true,
	},
}
