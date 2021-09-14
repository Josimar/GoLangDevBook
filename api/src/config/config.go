package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco -> string of connection to database
	StringConexaoBanco = ""

	// Porta onde a API vai rodar
	Porta = 0

	// SecretKey -> key to use to check the token
	SecretKey []byte
)

// Carregar -> vai carregar as váriaveis do ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_BASE"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
