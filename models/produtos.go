package models

import (
	// "github.com/seu-usuario/crud_product_go/db"
	"github.com/seu-usuario/crud_product_go/db"
)

type Produto struct {
	ID int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

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
	defer db.Close()
	return produtos
}