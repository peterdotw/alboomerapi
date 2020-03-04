package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterdotw/alboomerapi/database"
	"github.com/peterdotw/alboomerapi/structs"
)

// AlbumPutHandler - Album PUT Handler
func AlbumPutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedAlbum structs.Album
	var existingAlbumID string
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updatedAlbum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	row, err := database.Dot.QueryRow(database.Db, "select-artist-id", updatedAlbum.ArtistName)
	err = row.Scan(&existingAlbumID)
	if err != nil {
		if err == sql.ErrNoRows {
			result, err := database.Dot.Exec(database.Db, "create-artist", updatedAlbum.ArtistName)
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			artistID, err := result.LastInsertId()
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			_, err = database.Dot.Exec(database.Db, "update-album", updatedAlbum.Name, artistID, updatedAlbum.ReleaseDate, updatedAlbum.Genre, params["id"])
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	updateResult, err := database.Dot.Exec(database.Db, "update-album", updatedAlbum.Name, existingAlbumID, updatedAlbum.ReleaseDate, updatedAlbum.Genre, params["id"])
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
