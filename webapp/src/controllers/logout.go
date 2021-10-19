package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout -> clean cookie
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", 302)
}
