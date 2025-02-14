package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPostgres() *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("PostgreSQL not responding:", err)
	}

	return db
}
