package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	Host     = os.Getenv("DB_HOST")
	Port     = os.Getenv("DB_PORT")
	Username = os.Getenv("DB_USERNAME")
	Password = os.Getenv("DB_PASSWORD")
	DBName   = os.Getenv("DB_NAME")
)

func ConnectDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, Username, Password, DBName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
