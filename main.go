package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"html/template"
	"net/http"
)

func conectaComBancoDeDados() *sql.DB{
	conexao := "user=postgres dbname=loja_alura password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	ID int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil{
		panic(err.Error())
	}

	//Cada produto do banco de dados vou chamar de p, assim que eu criar todos os meus produtos, 
	//eu vou inserir uma variável chamada produtos a nossa lista, nosso slice, passo ele embaixo.
	p := Produto{}
	produtos := []Produto{}

	//o for, linha a linha vai vendo o que cada produto é, for do select de produtos e usar uma função chamada next. Ou seja, eu quero a próximo à linha – leu a primeira eu quero a próxima linha, vou abrir parênteses, abro e fecho chaves também.
	for selectDeTodosOsProdutos.Next(){
		var id, quantidade int 
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
		    
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()

}

	
