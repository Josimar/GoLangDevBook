package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CreateToken -> create a token
func CreateToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["expiration"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["id"] = usuarioId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidarToken -> validate the toke
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	//fmt.Println(token)

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

// ExtrairUsuarioID -> Extract user id
func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["id"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("invalid token")
}

// extrairToken -> Separate the token of Bearer
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Bearer token
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

// retornarChaveDeVerificacao -> Return key of token
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method unexpected! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
