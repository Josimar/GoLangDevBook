package models

import (
	"errors"
	"strings"
	"time"
)

// Usuario -> estrutura da tabela usuário
type Usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Preparar -> preparar todos os campos recebidos
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

// validar -> validação dos campos
func (usuario *Usuario) validar() error {
	if usuario.Name == "" {
		return errors.New("O nome é obrigatório e não pode ficar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode ficar em branco")
	}
	if usuario.Password == "" {
		return errors.New("A senha é obrigatória e não pode ficar em branco")
	}

	return nil
}

// formatar -> formatação dos campos
func (usuario *Usuario) formatar() {
	usuario.Name = strings.TrimSpace(usuario.Name)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
