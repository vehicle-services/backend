package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"technician/config"

	_ "github.com/lib/pq"
)

var dbInstance *sql.DB
var dbOnce sync.Once

func NewPostgresStorage(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.Port, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func InitPostgresDB() {
	db, err := NewPostgresStorage(config.Envs)
	if err != nil {
		log.Fatal("Could not get PostgreSQL Storage")
	}
	dbInstance = db
}

func GetPostgresDB() *sql.DB {
	dbOnce.Do(InitPostgresDB)
	return dbInstance
}
