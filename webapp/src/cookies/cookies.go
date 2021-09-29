package cookies

import (
	"github.com/gorilla/securecookie"
	"net/http"
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
