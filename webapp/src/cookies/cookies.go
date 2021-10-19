package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
	"webapp/src/config"
)

var s *securecookie.SecureCookie

// Configurar -> configurar cookies
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// Salvar -> salva as infos do cookie
func Salvar(w http.ResponseWriter, Id, token string) error {
	dados := map[string]string{
		"id":    Id,
		"token": token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Ler -> check if cookie exists and return value in cookies
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}

// Deletar -> clean the cookie
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
