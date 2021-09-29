package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger -> log da web
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar -> check is cookie exists
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// valores, erro := cookies.Ler(r)
		// fmt.Println(valores, erro)

		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		proximaFuncao(w, r)
	}
}
