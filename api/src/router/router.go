package router

import "github.com/gorilla/mux"

// Gerar -> retorna uma router com suas rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
