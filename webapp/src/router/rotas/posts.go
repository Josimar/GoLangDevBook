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
}
