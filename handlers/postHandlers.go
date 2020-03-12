package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/peterdotw/alboomerapi/database"
	"github.com/peterdotw/alboomerapi/structs"
)

// AlbumPostHandler - Album POST Handler
func AlbumPostHandler(w http.ResponseWriter, r *http.Request) {
	var newAlbum structs.Album
	var existingAlbumID string

	err := json.NewDecoder(r.Body).Decode(&newAlbum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	row, err := database.Dot.QueryRow(database.Db, "select-artist-id", newAlbum.ArtistName)
	err = row.Scan(&existingAlbumID)
	if err != nil {
		if err == sql.ErrNoRows {
			result, _ := database.Dot.Exec(database.Db, "create-artist", newAlbum.ArtistName)

			artistID, _ := result.LastInsertId()

			_, err = database.Dot.Exec(database.Db, "create-album", newAlbum.Name, artistID, newAlbum.ReleaseDate, newAlbum.Genre)
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}
	}

	database.Dot.Exec(database.Db, "create-album", newAlbum.Name, existingAlbumID, newAlbum.ReleaseDate, newAlbum.Genre)

	w.WriteHeader(http.StatusCreated)
}
