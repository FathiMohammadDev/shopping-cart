package main

import (
	"database/sql"
	"log"

	"github.com/FathiMohammadDev/shopping-cart/cmd/api"
	"github.com/FathiMohammadDev/shopping-cart/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	server.Run()

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB contected successfuly")
}
