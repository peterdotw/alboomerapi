package handlers

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestAlbumPostHandler(t *testing.T) {
	initDatabases()
	payload := []byte(`{"album_name":"Thirst","artist_name":"SebastiAn","release_date":"2019-11-08","genre":"Electronic"}`)

	req, err := http.NewRequest("POST", "/api/v1/album", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestAlbumPostHandlerWhenArtistExist(t *testing.T) {
	initDatabases()
	payload := []byte(`{"album_name":"Lonerism","artist_name":"Tame Impala","release_date":"2012-10-05","genre":"Electronic"}`)

	req, err := http.NewRequest("POST", "/api/v1/album", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestAlbumPostHandlerWithBadPayload(t *testing.T) {
	initDatabases()
	badPayload := []byte("example bad payload")

	req, err := http.NewRequest("POST", "/api/v1/album", bytes.NewBuffer(badPayload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}
