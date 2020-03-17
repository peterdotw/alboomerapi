package main

import (
	"log"
	"net/http"

	"github.com/peterdotw/alboomerapi/routes"
	"github.com/peterdotw/alboomerapi/utils"

	"github.com/rs/cors"
)

// PORT : Port on which app is running
const PORT = ":3000"

func main() {
	routesHandler := utils.RequestLoggerMiddleware(cors.Default().Handler(routes.CreateRouter()))
	log.Println("Server running on port", PORT)
	log.Fatal(http.ListenAndServe(PORT, routesHandler))
}
