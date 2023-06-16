package db

import (_ "github.com/lib/pq"
"database/sql")

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=loja_alura password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}