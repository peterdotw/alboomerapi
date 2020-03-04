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
	_, err = dot.Exec(db, "create-artists-table")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Exec(db, "create-albums-table")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dot.Exec(db, "create-tracks-table")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// InitDotSQL - Initialize DotSQL
func InitDotSQL() *dotsql.DotSql {
	dotAlbums, err := dotsql.LoadFromFile("database/tables/albums.sql")
	if err != nil {
		log.Fatal(err)
	}
	dotArtists, err := dotsql.LoadFromFile("database/tables/artists.sql")
	if err != nil {
		log.Fatal(err)
	}
	dotTracks, err := dotsql.LoadFromFile("database/tables/tracks.sql")
	if err != nil {
		log.Fatal(err)
	}

	dot := dotsql.Merge(dotAlbums, dotArtists, dotTracks)

	return dot
}

//Db - Database initialized
var Db = InitDB()

//Dot - DotSQL initialized
var Dot = InitDotSQL()
