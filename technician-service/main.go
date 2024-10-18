package main

import (
	"database/sql"
	"log"
	"technician/db"
	"technician/server"
)

func main() {
	dbase := db.GetPostgresDB()
	initStorage(dbase)
	
	server := server.NewAPIServer(":8000", dbase)
	err := server.Run()
	if err != nil {
		log.Fatal("Could not Run server")
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to DB")
}