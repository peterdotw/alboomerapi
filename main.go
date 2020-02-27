package main

import (
	"log"
	"net/http"

	"github.com/peterdotw/alboomerapi/routes"
)

// PORT : Port on which app is running
const PORT = ":3000"

func main() {
	routesHandler := routes.MakeHTTPHandler()
	log.Println("Server running on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, routesHandler))
}
