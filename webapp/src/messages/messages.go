package messages

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroApi -> error constants
type ErroApi struct {
	Erro string `json:"erro"`
}

// JSON -> retorna uma resposta em formato JSON
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// TratarStatusCodeDeErro -> Tratar status do erro
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroApi
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
