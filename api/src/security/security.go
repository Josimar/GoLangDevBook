package security

import "golang.org/x/crypto/bcrypt"

// Hash -> recebe uma string e criptografa
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerificarSenha -> verifica a senha do usuário
func VerificarSenha(senhaString, senhaComHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaString), []byte(senhaComHash))
}
