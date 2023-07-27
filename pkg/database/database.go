package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	Conn *sql.DB
}

func NewMySQLDB(dataSourceName string) *MySQLDB {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	fmt.Println("Connected to the database!")

	return &MySQLDB{
		Conn: db,
	}
}