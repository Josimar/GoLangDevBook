package rotas

import "net/http"

// Rota -> representa todas as rotas da API
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}
