package middlewares

import (
	"api/src/authentication"
	"api/src/messages"
	"log"
	"net/http"
)

// Logger -> print information on the terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		next(w, r)
	}
}

// Autenticar -> middleware to authentic token
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := authentication.ValidarToken(r); erro != nil {
			messages.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
