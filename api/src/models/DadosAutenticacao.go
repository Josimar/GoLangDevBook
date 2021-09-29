package models

// DadosAutenticacao -> contem dados da autenticação
type DadosAutenticacao struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
