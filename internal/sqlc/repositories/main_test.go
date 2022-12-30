package db

import (
	"database/sql"
	"log"
	"os"
	"template-go/internal/config"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../../configs")
	if err != nil {
		log.Fatal("cannot load config env:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
