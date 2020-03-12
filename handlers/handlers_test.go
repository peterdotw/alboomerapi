package handlers

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/peterdotw/alboomerapi/database"
)

func mockRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/albums", AlbumsGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album", AlbumPostHandler).Methods("POST")
	router.HandleFunc("/api/v1/album/{id:[0-9]+}", AlbumGetHandler).Methods("GET")
	router.HandleFunc("/api/v1/album/{id:[0-9]+}", AlbumPutHandler).Methods("PUT")
	router.HandleFunc("/api/v1/album/{id:[0-9]+}", AlbumDeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/artists", ArtistsGetHandler).Methods("GET")
	return router
}

var testRouter = mockRouter()

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	testRouter.ServeHTTP(recorder, req)

	return recorder
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, r *httptest.ResponseRecorder, expected string) {
	if strings.TrimSpace(r.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", r.Body.String(), expected)
	}
}

func initDatabases() {
	database.RedisConnection.Do("FLUSHALL")
	database.Db.Exec("DROP TABLE tracks;")
	database.Db.Exec("DROP TABLE albums;")
	database.Db.Exec("DROP TABLE artists;")
	database.Dot.Exec(database.Db, "create-artists-table")
	database.Dot.Exec(database.Db, "create-albums-table")
	database.Dot.Exec(database.Db, "create-tracks-table")
	database.Dot.Exec(database.Db, "create-artist", "Grimes")
	database.Dot.Exec(database.Db, "create-artist", "Mac Demarco")
	database.Dot.Exec(database.Db, "create-artist", "Tame Impala")
	database.Dot.Exec(database.Db, "create-album", "Miss Anthropocene", 1, "2020-02-21", "Electronic")
	database.Dot.Exec(database.Db, "create-album", "Salad Days", 2, "2014-04-01", "Rock")
	database.Dot.Exec(database.Db, "create-album", "The Slow Rush", 3, "2020-02-14", "Rock")
}

func clearDatabases() {
	database.RedisConnection.Do("FLUSHALL")
	database.Db.Exec("DROP TABLE tracks;")
	database.Db.Exec("DROP TABLE albums;")
	database.Db.Exec("DROP TABLE artists;")
}

func addAlbum(ID int, name, releaseDate, genre string) {
	database.Dot.Exec(database.Db, "create-album", name, ID, releaseDate, genre)
}

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

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"albums":[{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"},{"album_id":2,"album_name":"Salad Days","artist_name":"Mac Demarco","release_date":"2014-04-01","genre":"Rock"},{"album_id":3,"album_name":"The Slow Rush","artist_name":"Tame Impala","release_date":"2020-02-14","genre":"Rock"}]}`

	checkResponseBody(t, response, expected)
}

func TestGetNonExistentAlbums(t *testing.T) {
	clearDatabases()
	req, err := http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"albums":null}`

	checkResponseBody(t, response, expected)
}

func TestGetCachedAlbums(t *testing.T) {
	initDatabases()
	req, err := http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"albums":[{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"},{"album_id":2,"album_name":"Salad Days","artist_name":"Mac Demarco","release_date":"2014-04-01","genre":"Rock"},{"album_id":3,"album_name":"The Slow Rush","artist_name":"Tame Impala","release_date":"2020-02-14","genre":"Rock"}]}`

	checkResponseBody(t, response, expected)

	req, err = http.NewRequest("GET", "/api/v1/albums", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected = `{"albums":[{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"},{"album_id":2,"album_name":"Salad Days","artist_name":"Mac Demarco","release_date":"2014-04-01","genre":"Rock"},{"album_id":3,"album_name":"The Slow Rush","artist_name":"Tame Impala","release_date":"2020-02-14","genre":"Rock"}]}`

	checkResponseBody(t, response, expected)
}

func TestArtistsGetHandler(t *testing.T) {
	initDatabases()
	req, err := http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"artists":[{"artist_id":1,"artist_name":"Grimes"},{"artist_id":2,"artist_name":"Mac Demarco"},{"artist_id":3,"artist_name":"Tame Impala"}]}`

	checkResponseBody(t, response, expected)
}

func TestGetNonExistentArtists(t *testing.T) {
	clearDatabases()
	req, err := http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"artists":null}`

	checkResponseBody(t, response, expected)
}

func TestGetCachedArtists(t *testing.T) {
	initDatabases()
	req, err := http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"artists":[{"artist_id":1,"artist_name":"Grimes"},{"artist_id":2,"artist_name":"Mac Demarco"},{"artist_id":3,"artist_name":"Tame Impala"}]}`

	checkResponseBody(t, response, expected)

	req, err = http.NewRequest("GET", "/api/v1/artists", nil)
	if err != nil {
		log.Fatalln(err)
	}

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected = `{"artists":[{"artist_id":1,"artist_name":"Grimes"},{"artist_id":2,"artist_name":"Mac Demarco"},{"artist_id":3,"artist_name":"Tame Impala"}]}`

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

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{"album_id":1,"album_name":"Miss Anthropocene","artist_name":"Grimes","release_date":"2020-02-21","genre":"Electronic"}`

	checkResponseBody(t, response, expected)
}

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
	badPayload := []byte("YOU JUST GOT PRANKED BRO")

	req, err := http.NewRequest("POST", "/api/v1/album", bytes.NewBuffer(badPayload))
	if err != nil {
		log.Fatalln(err)
	}
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

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

func init() {
	initDatabases()
	log.Println("Tests started")
}
