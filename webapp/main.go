package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

// init -> gerar random key
/*
func init(){
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)
}
*/

// main -> Main Function
func main() {

	config.Carregar()

	cookies.Configurar()

	utils.CarregarTemplates()

	r := router.Gerar()

	fmt.Printf("Started WebApp! (%d)", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
