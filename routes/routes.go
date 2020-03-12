package routes

import (
	"net/http"

	"github.com/peterdotw/alboomerapi/handlers"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte("<h1>Example REST Api</h1><p>Example REST Api written entirely in Go."))
}

// CreateRouter - Router
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/api/v1/albums", handlers.AlbumsGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album", handlers.AlbumPostHandler).Methods("POST")
	router.HandleFunc("/api/v1/album/{id:[0-9]+}", handlers.AlbumGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album/{id:[0-9]+}", handlers.AlbumPutHandler).Methods("PUT")
	router.HandleFunc("/api/v1/album/{id:[0-9]+}", handlers.AlbumDeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/artists", handlers.ArtistsGetHandler).Methods("GET")
	return router
}
