package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Could not ping PostgreSQL: %v", err)
	}

	// Create tables if they don't exist
	createTablesSQL := `
	CREATE TABLE IF NOT EXISTS url_clicks (
		id SERIAL PRIMARY KEY,
		short_key VARCHAR(10) NOT NULL,
		click_time TIMESTAMP NOT NULL DEFAULT NOW(),
		ip_address VARCHAR(50),
		country VARCHAR(100),
		city VARCHAR(100),
		region VARCHAR(100)
	);
	`

	_, err = DB.Exec(createTablesSQL)
	if err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	log.Println("Connected to PostgreSQL and initialized schema")
}
