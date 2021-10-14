package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPosts = []Rota{
	{
		Uri:                "/posts",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPosts,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}/like",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPost,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/posts/{postId}/dislike",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DesCurtirPost,
		RequerAutenticacao: true,
	},
}
