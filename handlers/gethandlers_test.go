package handlers

import (
	"log"
	"net/http"
	"testing"
)

func TestGetNonExistentAlbum(t *testing.T) {
	initDatabases()

	req, err := http.NewRequest("GET", "/api/v1/album/7777", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestAlbumsGetHandler(t *testing.T) {
	initDatabases()

	req, err := http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	expected := `{"albums":[{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"},{"album_id":2,"album_name":"Salad Days","artist_name":"Mac Demarco","release_date":"2014-04-01","genre":"Rock"},{"album_id":3,"album_name":"The Slow Rush","artist_name":"Tame Impala","release_date":"2020-02-14","genre":"Rock"}]}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}

func TestGetNonExistentAlbums(t *testing.T) {
	clearDatabases()

	req, err := http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	expected := `{"albums":null}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}

func TestGetCachedAlbums(t *testing.T) {
	initDatabases()

	req, err := http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	req, err = http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)
	expected := `{"albums":[{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"},{"album_id":2,"album_name":"Salad Days","artist_name":"Mac Demarco","release_date":"2014-04-01","genre":"Rock"},{"album_id":3,"album_name":"The Slow Rush","artist_name":"Tame Impala","release_date":"2020-02-14","genre":"Rock"}]}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}

func TestArtistsGetHandler(t *testing.T) {
	initDatabases()

	req, err := http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	expected := `{"artists":[{"artist_id":1,"artist_name":"Grimes"},{"artist_id":2,"artist_name":"Mac Demarco"},{"artist_id":3,"artist_name":"Tame Impala"}]}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}

func TestGetNonExistentArtists(t *testing.T) {
	clearDatabases()

	req, err := http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	expected := `{"artists":null}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}

func TestGetCachedArtists(t *testing.T) {
	initDatabases()

	req, err := http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	req, err = http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)
	expected := `{"artists":[{"artist_id":1,"artist_name":"Grimes"},{"artist_id":2,"artist_name":"Mac Demarco"},{"artist_id":3,"artist_name":"Tame Impala"}]}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}

func TestAlbumGetHandler(t *testing.T) {
	initDatabases()
	addAlbum(1, "Grimes", "2020-02-21", "Electronic")

	req, err := http.NewRequest("GET", "/api/v1/album/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	expected := `{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"}`

	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, response, expected)
}
