package main

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/seu-usuario/crud_product_go/models"
)



var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}

	
