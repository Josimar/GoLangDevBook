package controllers

import "net/http"

// CriarUsuario
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Controllers -> CriarUsuario -> Criando usuário"))
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
