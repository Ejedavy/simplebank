package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testStore *Store

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Could not establish the database connection")
	}
	testQueries = New(conn)
	testStore, err = NewStore(testQueries, conn)
	if err != nil {
		log.Fatal("Could not create a transaction wrapper")
	}
	os.Exit(m.Run())
}
