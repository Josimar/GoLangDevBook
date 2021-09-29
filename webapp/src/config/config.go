package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// Porta -> Connection Port
	Porta = 0

	// ApiUrl -> url da API
	ApiUrl = ""

	// HashKey -> utilizado para autenticar cookie
	HashKey []byte

	// BlockKey -> utilizado para criptografar o cookie
	BlockKey []byte
)

// Carregar -> vai carregar as váriaveis do ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 8001
	}

	ApiUrl = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
