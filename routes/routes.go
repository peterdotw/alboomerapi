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
	Artist      string `json:"artist"`
	ReleaseDate string `json:"release_date"` //time.Time
	Genre       string `json:"genre"`
	//Tracks []Track  `json:"tracks"`
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

func albumsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		albumsGetHandler(w, r)
	case "POST":
		albumsPostHandler(w, r)
	default:
		w.WriteHeader(http.StatusUnauthorized)
	}
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

func albumHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		albumGetHandler(w, r)
	case "POST":
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusUnauthorized)
	}
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

// MakeHTTPHandler - Handler for routes
func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/api/v1/albums", albumsHandler)
	router.HandleFunc("/api/v1/album/{id}", albumHandler)
	return router
}
