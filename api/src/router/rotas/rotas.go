package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Rota -> representa todas as rotas da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Configurar -> Coloca todas as rotas dentro do Router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		r.HandleFunc(rota.Uri, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
