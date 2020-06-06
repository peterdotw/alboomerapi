package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/peterdotw/alboomerapi/database"
	"github.com/peterdotw/alboomerapi/structs"
)

// AlbumsGetHandler - Albums GET Handler
func AlbumsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var album structs.Album
	var albums structs.Albums

	getAlbumsFromRedis, err := redis.Bytes(database.RedisConnection.Do("GET", "/albums"))
	if err != nil {
		rows, _ := database.Dot.Query(database.Db, "select-albums")
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&album.ID, &album.Name, &album.ArtistName, &album.ReleaseDate, &album.Genre)
			albums.Albums = append(albums.Albums, album)
		}

		albumsBytes, _ := json.Marshal(albums)
		json.Unmarshal(albumsBytes, &albums)
		json.NewEncoder(w).Encode(albums)

		database.RedisConnection.Do("SETEX", "/albums", 86400, albumsBytes)

		return
	}

	json.Unmarshal(getAlbumsFromRedis, &albums)
	json.NewEncoder(w).Encode(albums)
}

// AlbumGetHandler - Album GET Handler
func AlbumGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var album structs.Album

	getAlbumFromRedis, err := redis.Bytes(database.RedisConnection.Do("GET", "/album/"+params["id"]))
	if err != nil {
		row, _ := database.Dot.QueryRow(database.Db, "select-album", params["id"])
		err := row.Scan(&album.ID, &album.Name, &album.ArtistName, &album.ReleaseDate, &album.Genre)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		albumBytes, _ := json.Marshal(album)
		json.NewEncoder(w).Encode(album)

		database.RedisConnection.Do("SETEX", "/album/"+params["id"], 86400, albumBytes)

		return
	}

	json.Unmarshal(getAlbumFromRedis, &album)
	json.NewEncoder(w).Encode(album)
}

// ArtistsGetHandler - Artists GET Handler
func ArtistsGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var artist structs.Artist
	var artists structs.Artists

	getArtistsFromRedis, err := redis.Bytes(database.RedisConnection.Do("GET", "/artists"))
	if err != nil {
		rows, _ := database.Dot.Query(database.Db, "select-artists")
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&artist.ID, &artist.ArtistName)
			artists.Artists = append(artists.Artists, artist)
		}

		artistsBytes, _ := json.Marshal(artists)
		json.Unmarshal(artistsBytes, &artists)
		json.NewEncoder(w).Encode(artists)

		database.RedisConnection.Do("SETEX", "/artists", 86400, artistsBytes)

		return
	}

	json.Unmarshal(getArtistsFromRedis, &artists)
	json.NewEncoder(w).Encode(artists)
}
