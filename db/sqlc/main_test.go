package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	testDB, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("can`t  connect to db: %v", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
