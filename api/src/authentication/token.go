package authentication

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken -> create a token
func CreateToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["expiration"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["id"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("DevBookSecret"))
}
