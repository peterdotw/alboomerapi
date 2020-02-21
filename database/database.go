package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql" // MySQL Driver
)

// InitDB - Initialize Database connection
func InitDB() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	name := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, address, name))
	if err != nil {
		log.Fatal(err)
	}
	dot := InitDotSQL()
	res, err := dot.Exec(db, "create-albums-table")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.RowsAffected())

	return db
}

// InitDotSQL - Initialize DotSQL
func InitDotSQL() *dotsql.DotSql {
	dot, err := dotsql.LoadFromFile("database/tables/albums.sql")
	if err != nil {
		log.Fatal(err)
	}

	return dot
}
