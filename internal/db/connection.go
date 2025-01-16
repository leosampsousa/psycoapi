package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil 

//TODO: criar variaveis de ambiente
func newDB(url string) *sql.DB {
	connStr := url
    db, err := sql.Open("postgres", connStr)

    if err != nil {
        log.Fatal(err)
    }

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return db
}

func GetInstance(url string) *sql.DB {
	if db != nil {
		return db
	}

	db = newDB(url)
	return db
}