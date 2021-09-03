package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CriarUsuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	repositorio := repository.NewUserRepository(db)
	usuarioId, erro := repositorio.Create(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", usuarioId)))
}

// BuscarUsuarios -> busca todos usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> BuscarUsuario -> Buscar usuários"))
}

// BuscarUsuario -> busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> BuscarUsuario -> Buscar um usuário"))
}

// AtualizarUsuario -> atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> AtualizarUsuario -> Atualizar um usuário"))
}

// ExcluirUsuario -> exclui um usuário
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> ExcluirUsuario -> Excluir um usuário"))
}
