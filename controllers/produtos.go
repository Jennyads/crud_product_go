package controllers

import (
	"net/http"
	"text/template"

	"github.com/seu-usuario/crud_product_go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}