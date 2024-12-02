package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	var err error
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),
	}
	DB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Print("Error connecting to the database: ", err)
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Print("Error pinging the database: ", err)
		log.Fatal(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			price INT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Error creating products table: %v", err)
	}

	log.Print("Connected to the database")
	return DB, nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}
