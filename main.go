package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/peterdotw/alboomerapi/routes"

	_ "github.com/go-sql-driver/mysql"
)

// PORT : Port on which app is running
const PORT = ":3000"

func main() {
	routesHandler := routes.MakeHTTPHandler()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	name := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, address, name))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Server running on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, routesHandler))
}
