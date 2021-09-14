package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
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
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// validar -> validação dos campos
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Name == "" {
		return errors.New("O nome é obrigatório e não pode ficar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode ficar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Password == "" {
		return errors.New("A senha é obrigatória e não pode ficar em branco")
	}

	return nil
}

// formatar -> formatação dos campos
func (usuario *Usuario) formatar(etapa string) error {
	usuario.Name = strings.TrimSpace(usuario.Name)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Password)
		if erro != nil {
			return erro
		}
		usuario.Password = string(senhaComHash)
	}

	return nil
}
