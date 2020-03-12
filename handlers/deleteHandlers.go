package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterdotw/alboomerapi/database"
)

// AlbumDeleteHandler - Album DELETE Handler
func AlbumDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	res, _ := database.Dot.Exec(database.Db, "delete-album", params["id"])
	rowsCount, _ := res.RowsAffected()
	if rowsCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	database.RedisConnection.Do("DEL", "/album/"+params["id"])

	w.WriteHeader(http.StatusOK)
}
