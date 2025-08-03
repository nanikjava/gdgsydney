package main

import (
	"fmt"
	"gdgsydney/db"
	"gdgsydney/routes"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	routes.SetupRoutes(db)

	fmt.Println("Server is starting on port 7666...")
	log.Fatal(http.ListenAndServe(":7666", nil))
}
