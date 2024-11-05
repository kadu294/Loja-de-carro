package db

import (
	"database/sql"
)

func ConectaComBancoDeDados() *sql.DB {
	conexao := "host=previg-dev02 port=5434 dbname=dbprevig04 user=carlos_dev password=carlos_dev sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
