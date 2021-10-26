package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaCadastroUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/buscar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuario/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuario/{usuarioId}/unfollow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuario/{usuarioId}/follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilUsuarioLogado,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDeEdicaoUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/atualizar-senha",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeAtualizacaoUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-usuario",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/editar-usuario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.EditarUsuario,
		RequerAutenticacao: true,
	},
}
