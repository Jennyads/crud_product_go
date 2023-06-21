package main

import (
	"net/http"

	"github.com/seu-usuario/crud_product_go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}



	
