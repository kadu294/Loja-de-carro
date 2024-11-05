package main

import (
	"net/http"

	"github.com/kadu294/Loja-de-carro/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
