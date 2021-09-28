package router

import (
	"github.com/gorilla/mux"
	"webapp/src/router/rotas"
)

// Gerar -> retorna um router com todas as rotas geradas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}
