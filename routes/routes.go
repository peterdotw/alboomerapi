package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/peterdotw/alboomerapi/database"

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

var db = database.InitDB()
var dot = database.InitDotSQL()

func makeExampleDB() Albums {
	var album Album
	var albums Albums
	rows, err := dot.Query(db, "select-albums")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&album.ID, &album.Name, &album.Artist, &album.ReleaseDate, &album.Genre)
		if err != nil {
			log.Fatal(err)
		}
		albums.Albums = append(albums.Albums, album)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	allAlbums, _ := json.Marshal(albums)
	json.Unmarshal(allAlbums, &albums)
	return albums
}

var initAlbums Albums = makeExampleDB()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte("<h1>Example REST Api</h1><p>Example REST Api written entirely in Go without any external modules except for go-sql-driver."))
}

func albumsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var album Album
	var albums Albums
	rows, err := dot.Query(db, "select-albums")
	if err != nil {
		json.NewEncoder(w).Encode(Albums{})
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&album.ID, &album.Name, &album.Artist, &album.ReleaseDate, &album.Genre)
		if err != nil {
			log.Fatal(err)
		}
		albums.Albums = append(albums.Albums, album)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	allAlbums, _ := json.Marshal(albums)
	json.Unmarshal(allAlbums, &albums)
	json.NewEncoder(w).Encode(albums)
}

func albumsPostHandler(w http.ResponseWriter, r *http.Request) {
	var newAlbum Album
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.Unmarshal(body, &newAlbum)
	res, err := dot.Exec(db, "create-album", newAlbum.Name, newAlbum.Artist, newAlbum.ReleaseDate, newAlbum.Genre)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	initAlbums.Albums = append(initAlbums.Albums, newAlbum)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newAlbum)
}

func albumGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var album Album
	row, _ := dot.QueryRow(db, "select-album", params["id"])
	err := row.Scan(&album.ID, &album.Name, &album.Artist, &album.ReleaseDate, &album.Genre)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(album)
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
func MakeHTTPHandler() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/api/v1/albums", albumsGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/albums", albumsPostHandler).Methods("POST")
	router.HandleFunc("/api/v1/album/{id}", albumGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album/{id}", albumPutHandler).Methods("PUT")
	router.HandleFunc("/api/v1/album/{id}", albumDeleteHandler).Methods("DELETE")
	return router
}