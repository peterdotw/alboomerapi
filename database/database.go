package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/gchaincl/dotsql"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", username+":"+password+"@/"+name)
	if err != nil {
		log.Fatal(err)
	}
	dot := initDotSQL()
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

func initDotSQL() *dotsql.DotSql {
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
var Db = initDB()

//Dot - DotSQL initialized
var Dot = initDotSQL()
