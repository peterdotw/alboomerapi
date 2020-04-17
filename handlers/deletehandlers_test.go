package handlers

import (
	"log"
	"net/http"
	"testing"
)

func TestAlbumDeleteHandler(t *testing.T) {
	initDatabases()
	req, err := http.NewRequest("GET", "/api/v1/album/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("DELETE", "/api/v1/album/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, err = http.NewRequest("GET", "/api/v1/album/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	req, err = http.NewRequest("DELETE", "/api/v1/album/1", nil)
	if err != nil {
		log.Fatalln(err)
	}
	response = executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}
