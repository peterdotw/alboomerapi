package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/gchaincl/dotsql"
	"github.com/joho/godotenv"
)

// InitDB - Initialize Database connection
func InitDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	instance := os.Getenv("INSTANCE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	cfg := mysql.Cfg(instance, username, password)
	cfg.DBName = os.Getenv("DB_NAME")
	db, err := mysql.DialCfg(cfg)
	if err != nil {
		log.Fatal(err)
	}
	dot := InitDotSQL()
	_, err = dot.Exec(db, "create-albums-table")
	if err != nil {
		log.Fatal(err)
	}

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
