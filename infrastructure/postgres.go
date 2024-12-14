package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var POSTGRES_DB *sql.DB
var POSTGRES_CONNECTION_STRING string

func init() {
	InitializePostgresSQL()
}

func InitializePostgresSQL() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Couldn't connect to database")
	}
	USER := os.Getenv("USER")
	HOST := os.Getenv("HOST")
	DBNAME := os.Getenv("DBNAME")
	PASSWORD := os.Getenv("PASSWORD")
	PORT := os.Getenv("PORT")

	POSTGRES_CONNECTION_STRING = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", HOST, USER, PASSWORD, DBNAME, PORT)

	POSTGRES_DB, err = sql.Open("pgx", POSTGRES_CONNECTION_STRING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to postgres db:%v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Connected to database")
	}

	POSTGRES_DB.SetMaxIdleConns(10)

}
