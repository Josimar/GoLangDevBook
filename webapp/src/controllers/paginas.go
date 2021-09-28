package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin -> render the login screen
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Login screen"))
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaCadastroUsuario -> render register page
func CarregarPaginaCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
