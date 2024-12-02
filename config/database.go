package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() (*sql.DB, error) {
	var err error

	dsn := os.Getenv("DB_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	DB, err = sql.Open("mysql", dsn)
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
            buyer_name VARCHAR(255) NOT NULL,
            store_name VARCHAR(255) NOT NULL,
            item_name VARCHAR(255) NOT NULL,
            item_qty INT NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		log.Fatalf("Error creating orders table: %v", err)
	}

	log.Print("Connected to the database")
	return DB, nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}
