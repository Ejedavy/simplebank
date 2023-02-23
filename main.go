package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"simple_bank/api"
	db "simple_bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("Could not start database", err)
	}
	queries := db.New(conn)

	store, err := db.NewStore(queries, conn)
	if err != nil {
		log.Fatalln("Could not create queries", err)
	}

	server := api.NewServer(store)
	fmt.Println("Starting the server at", address)
	server.Start(address)
}
