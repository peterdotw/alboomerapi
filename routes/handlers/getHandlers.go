package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterdotw/alboomerapi/database"
	"github.com/peterdotw/alboomerapi/structs"
)

// AlbumsGetHandler - Albums GET Handler
func AlbumsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var album structs.Album
	var albums structs.Albums

	rows, err := database.Dot.Query(database.Db, "select-albums")
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

// AlbumGetHandler - Album GET Handler
func AlbumGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var album structs.Album

	row, _ := database.Dot.QueryRow(database.Db, "select-album", params["id"])
	err := row.Scan(&album.ID, &album.Name, &album.ArtistName, &album.ReleaseDate, &album.Genre)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(album)
}

// ArtistsGetHandler - Artists GET Handler
func ArtistsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var artist structs.Artist
	var artists structs.Artists

	rows, err := database.Dot.Query(database.Db, "select-artists")
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
