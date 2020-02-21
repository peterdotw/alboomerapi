package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterdotw/alboomerapi/routes"
)

// PORT : Port on which app is running
const PORT = ":3000"

func main() {
	routesHandler := routes.MakeHTTPHandler()
	// Kurła
	log.Println("Server running on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, routesHandler))
}
