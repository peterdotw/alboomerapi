package handlers

import (
	"encoding/json"
	"log"
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

		albumsBytes, _ := json.Marshal(albums)
		json.Unmarshal(albumsBytes, &albums)
		json.NewEncoder(w).Encode(albums)

		_, err = database.RedisConnection.Do("SETEX", "/albums", 86400, albumsBytes)
		if err != nil {
			log.Panic(err)
		}

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

		_, err = database.RedisConnection.Do("SETEX", "/album/"+params["id"], 86400, albumBytes)
		if err != nil {
			log.Panic(err)
		}

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
		rows, err := database.Dot.Query(database.Db, "select-artists")
		if err != nil {
			json.NewEncoder(w).Encode(structs.Artists{})
			return
		}
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

		artistsBytes, _ := json.Marshal(artists)
		json.Unmarshal(artistsBytes, &artists)
		json.NewEncoder(w).Encode(artists)

		_, err = database.RedisConnection.Do("SETEX", "/artists", 86400, artistsBytes)
		if err != nil {
			log.Panic(err)
		}

		return
	}

	json.Unmarshal(getArtistsFromRedis, &artists)
	json.NewEncoder(w).Encode(artists)
}
