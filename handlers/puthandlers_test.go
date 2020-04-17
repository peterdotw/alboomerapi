package handlers

import (
	"bytes"
	"log"
	"net/http"
	"testing"
)

func TestAlbumPutHandler(t *testing.T) {
	initDatabases()
	payload := []byte(`{"album_name":"Flamagra","artist_name":"Flying Lotus","release_date":"2019-05-24","genre":"Electronic"}`)

	req, err := http.NewRequest("PUT", "/api/v1/album/2", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("GET", "/api/v1/album/2", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)

	expected := `{"album_id":2,"album_name":"Flamagra","artist_name":"Flying Lotus","release_date":"2019-05-24","genre":"Electronic"}`

	checkResponseBody(t, response, expected)
}

func TestAlbumPutHandlerWithBadPayload(t *testing.T) {
	initDatabases()
	badPayload := []byte("YOU JUST GOT PRANKED BRO")

	req, err := http.NewRequest("PUT", "/api/v1/album/2", bytes.NewBuffer(badPayload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestAlbumPutHandlerWithArtistAlreadyExisting(t *testing.T) {
	initDatabases()
	payload := []byte(`{"album_name":"Visions","artist_name":"Grimes","release_date":"2012-01-31","genre":"Electronic"}`)

	req, err := http.NewRequest("PUT", "/api/v1/album/2", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("GET", "/api/v1/album/2", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)

	expected := `{"album_id":2,"album_name":"Visions","artist_name":"Grimes","release_date":"2012-01-31","genre":"Electronic"}`

	checkResponseBody(t, response, expected)
}

func TestAlbumPutHandlerTwiceWithSamePayload(t *testing.T) {
	initDatabases()
	payload := []byte(`{"album_name":"Flamagra","artist_name":"Flying Lotus","release_date":"2019-05-24","genre":"Electronic"}`)

	req, err := http.NewRequest("PUT", "/api/v1/album/2", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("PUT", "/api/v1/album/2", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}
