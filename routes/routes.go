package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/peterdotw/alboomerapi/database"
	"github.com/peterdotw/alboomerapi/structs"

	"github.com/gorilla/mux"
)

var db = database.InitDB()
var dot = database.InitDotSQL()

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write([]byte("<h1>Example REST Api</h1><p>Example REST Api written entirely in Go."))
}

func albumsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var album structs.Album
	var albums structs.Albums

	rows, err := dot.Query(db, "select-albums")
	if err != nil {
		json.NewEncoder(w).Encode(structs.Albums{})
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&album.ID, &album.Name, &album.ArtistName, &album.ReleaseDate, &album.Genre)
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

func albumPostHandler(w http.ResponseWriter, r *http.Request) {
	var newAlbum structs.Album
	var artistExists = ""
	var existingArtistID = ""

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	json.Unmarshal(body, &newAlbum)

	row, err := dot.QueryRow(db, "artist-exists", newAlbum.ArtistName)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err = row.Scan(&artistExists)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if artistExists != "1" {
		result, err := dot.Exec(db, "create-artist", newAlbum.ArtistName)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		artistID, err := result.LastInsertId()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = dot.Exec(db, "create-album", newAlbum.Name, artistID, newAlbum.ReleaseDate, newAlbum.Genre)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	}

	row, err = dot.QueryRow(db, "select-artist-id", newAlbum.ArtistName)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	row.Scan(&existingArtistID)

	_, err = dot.Exec(db, "create-album", newAlbum.Name, existingArtistID, newAlbum.ReleaseDate, newAlbum.Genre)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func albumGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var album structs.Album

	row, _ := dot.QueryRow(db, "select-album", params["id"])
	err := row.Scan(&album.ID, &album.Name, &album.ArtistName, &album.ReleaseDate, &album.Genre)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(album)
}

func albumPutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedAlbum structs.Album
	var artistExists = ""
	var existingArtistID = ""
	params := mux.Vars(r)

	_ = json.NewDecoder(r.Body).Decode(&updatedAlbum)

	row, err := dot.QueryRow(db, "artist-exists", updatedAlbum.ArtistName)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err = row.Scan(&artistExists)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if artistExists != "1" {
		result, err := dot.Exec(db, "create-artist", updatedAlbum.ArtistName)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		artistID, err := result.LastInsertId()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err = dot.Exec(db, "update-album", updatedAlbum.Name, artistID, updatedAlbum.ReleaseDate, updatedAlbum.Genre)
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.WriteHeader(http.StatusCreated)
		return
	}

	row, err = dot.QueryRow(db, "select-artist-id", updatedAlbum.ArtistName)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	row.Scan(&existingArtistID)

	updateResult, err := dot.Exec(db, "update-album", updatedAlbum.Name, existingArtistID, updatedAlbum.ReleaseDate, updatedAlbum.Genre, params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	rowsAffected, _ := updateResult.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func albumDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	res, _ := dot.Exec(db, "delete-album", params["id"])
	rowsCount, _ := res.RowsAffected()
	if rowsCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func artistsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var artist structs.Artist
	var artists structs.Artists

	rows, err := dot.Query(db, "select-artists")
	if err != nil {
		json.NewEncoder(w).Encode(structs.Artists{})
		return
	}
	log.Println(rows.Columns())
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&artist.ID, &artist.ArtistName)
		if err != nil {
			log.Fatal(err)
		}

		artists.Artists = append(artists.Artists, artist)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	allArtists, _ := json.Marshal(artists)
	json.Unmarshal(allArtists, &artists)
	json.NewEncoder(w).Encode(artists)
}

// MakeHTTPHandler - Handler for routes
func MakeHTTPHandler() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/api/v1/albums", albumsGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album", albumPostHandler).Methods("POST")
	router.HandleFunc("/api/v1/album/{id}", albumGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album/{id}", albumPutHandler).Methods("PUT")
	router.HandleFunc("/api/v1/album/{id}", albumDeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/artists", artistsGetHandler).Methods("GET")
	return router
}
