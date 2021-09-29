package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/messages"
)

// CriarUsuario -> chamar API para criar usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	//name := r.FormValue("name")
	// fmt.Println(name)

	usuario, erro := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if erro != nil {
		messages.JSON(w, http.StatusBadRequest, messages.ErroApi{Erro: erro.Error()})
		return
	}

	// fmt.Println(usuario)
	// fmt.Println(bytes.NewBuffer(usuario))

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		messages.JSON(w, http.StatusInternalServerError, messages.ErroApi{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	// fmt.Println(response.Body)

	if response.StatusCode >= 400 {
		messages.TratarStatusCodeDeErro(w, response)
		return
	}

	messages.JSON(w, response.StatusCode, nil)
}
