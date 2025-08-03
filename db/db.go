package db

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//go:embed migrations/*.sql
var migrationFs embed.FS

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	//dbPath is for the name of the file
	//dbURL is the complete DSN which specify the drive ie: sqlite3
	dbPath := "gdgsydney.db"
	dbURL := "sqlite3://" + dbPath

	// specify migrations directory
	sourceDriver, err := iofs.New(migrationFs, "migrations")
	if err != nil {
		log.Println(err)
	}

	//open sqlite3 database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// run migration process
	m, err := migrate.NewWithSourceInstance("iofs", sourceDriver, dbURL)
	defer m.Close()

	if err != nil {
		log.Fatal("error instantiating migration object:", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("error migrating database: ", err)
	}

	return &Database{DB: db}, nil

}

func (d *Database) Close() error {
	return d.DB.Close()
}
