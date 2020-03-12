package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/peterdotw/alboomerapi/database"
	"github.com/peterdotw/alboomerapi/structs"
)

// AlbumPutHandler - Album PUT Handler
func AlbumPutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updatedAlbum structs.Album
	var existingArtistID string
	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&updatedAlbum)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	row, err := database.Dot.QueryRow(database.Db, "select-artist-id", updatedAlbum.ArtistName)
	err = row.Scan(&existingArtistID)
	if err != nil {
		if err == sql.ErrNoRows {
			result, _ := database.Dot.Exec(database.Db, "create-artist", updatedAlbum.ArtistName)

			artistID, _ := result.LastInsertId()

			_, err = database.Dot.Exec(database.Db, "update-album", updatedAlbum.Name, artistID, updatedAlbum.ReleaseDate, updatedAlbum.Genre, params["id"])
			if err != nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			database.RedisConnection.Do("DEL", "/album/"+params["id"])

			updatedAlbum.ID, _ = strconv.Atoi(params["id"])
			albumBytes, _ := json.Marshal(updatedAlbum)

			database.RedisConnection.Do("SETEX", "/album/"+params["id"], 86400, albumBytes)

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	updateResult, _ := database.Dot.Exec(database.Db, "update-album", updatedAlbum.Name, existingArtistID, updatedAlbum.ReleaseDate, updatedAlbum.Genre, params["id"])

	rowsAffected, _ := updateResult.RowsAffected()
	if rowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
