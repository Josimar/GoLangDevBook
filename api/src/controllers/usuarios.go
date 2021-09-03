package controllers

import (
	"api/src/banco"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
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
	repositorio.Create(usuario)
}

// BuscarUsuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> BuscarUsuario -> Buscar usuários"))
}

// BuscarUsuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> BuscarUsuario -> Buscar um usuário"))
}

// AtualizarUsuario
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> AtualizarUsuario -> Atualizar um usuário"))
}

// ExcluirUsuario
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> ExcluirUsuario -> Excluir um usuário"))
}
