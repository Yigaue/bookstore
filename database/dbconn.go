package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func DBConnect() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	config := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBHOST"),
		DBName: os.Getenv("DBNAME"),
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connection successful")
	return db
}