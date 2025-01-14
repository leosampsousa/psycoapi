package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil 

//TODO: criar variaveis de ambiente
func newDB() *sql.DB {
	connStr := "postgres://postgres:admin@localhost:5432/psycodb?sslmode=disable"
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

func GetInstance() *sql.DB {
	if db != nil {
		return db
	}

	db = newDB()
	return db
}