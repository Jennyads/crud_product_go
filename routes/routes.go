package routes

import (
	"net/http"

	"github.com/seu-usuario/crud_product_go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}