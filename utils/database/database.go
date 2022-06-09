package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	// postgres://shopvee:123456@product_db/shopvee?sslmode=disable
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("failed on opening connection postgres:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("failed on ping postgres server:", err)
		return nil, err
	}
	return db, nil
}
