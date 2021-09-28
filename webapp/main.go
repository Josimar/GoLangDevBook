package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

// main -> Main Function
func main() {

	config.Carregar()

	utils.CarregarTemplates()

	r := router.Gerar()

	fmt.Printf("Started WebApp! (%d)", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))

}
