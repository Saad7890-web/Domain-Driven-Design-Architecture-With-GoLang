package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Saad7890-web/internal/config"
	_ "github.com/lib/pq"
)


func NewPostgresDb(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err := db.Ping(); err != nil{
		log.Fatal("Failed to connect db:", err)
	}

	log.Println("Postgresql connected")
	return db
}