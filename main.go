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

// @title           Go Bank
// @version         1.0
// @description     A simple bank API
// @termsOfService  https://google.com

// @contact.name   David Oheji
// @contact.url    https://twitter.com/ejedavy
// @contact.email  ejeohejidavid@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
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
