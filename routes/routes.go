package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// Albums : Structure for albums
type Albums struct {
	Albums []Album `json:"albums"`
}

// Album : Structure for single album
type Album struct {
	ID          int    `json:"id"`
	Name        string `json:"album_name"`
	Artist      string `json:"artist_name"`
	ReleaseDate string `json:"release_date"` //time.Time
	Genre       string `json:"genre"`
}

func makeExampleDB() []byte {
	jsonFile, fileErr := os.Open("example-albums.json")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var albums Albums
	err := json.Unmarshal(byteValue, &albums)
	if err != nil {
		log.Fatal(err)
	}
	allAlbums, _ := json.Marshal(albums)
	return allAlbums
}

var initAlbums Albums
var converted = json.Unmarshal(makeExampleDB(), &initAlbums)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte("<h1>Example REST Api</h1><p>Example REST Api written entirely in Go without any external modules except for go-sql-driver."))
}

func albumsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(initAlbums)
}

func albumsPostHandler(w http.ResponseWriter, r *http.Request) {
	var newAlbum Album
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(body, &newAlbum)
	initAlbums.Albums = append(initAlbums.Albums, newAlbum)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAlbum)
}

func albumGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, album := range initAlbums.Albums {
		if strconv.Itoa(album.ID) == (params["id"]) {
			json.NewEncoder(w).Encode(album)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func albumPutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	updatedAlbumID, _ := strconv.Atoi(params["id"])
	for index, album := range initAlbums.Albums {
		if strconv.Itoa(album.ID) == params["id"] {
			initAlbums.Albums = append(initAlbums.Albums[:index], initAlbums.Albums[index+1:]...)
			var updatedAlbum Album
			_ = json.NewDecoder(r.Body).Decode(&updatedAlbum)
			updatedAlbum.ID = updatedAlbumID
			initAlbums.Albums = append(initAlbums.Albums, updatedAlbum)
			json.NewEncoder(w).Encode(updatedAlbum)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func albumDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deletedAlbumID, _ := strconv.Atoi(params["id"])
	for index, item := range initAlbums.Albums {
		if item.ID == deletedAlbumID {
			initAlbums.Albums = append(initAlbums.Albums[:index], initAlbums.Albums[index+1:]...)
			json.NewEncoder(w).Encode(initAlbums.Albums)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// MakeHTTPHandler - Handler for routes
func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/api/v1/albums", albumsGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/albums", albumsPostHandler).Methods("POST")
	router.HandleFunc("/api/v1/album/{id}", albumGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album/{id}", albumPutHandler).Methods("PUT")
	router.HandleFunc("/api/v1/album/{id}", albumDeleteHandler).Methods("DELETE")
	return router
}
