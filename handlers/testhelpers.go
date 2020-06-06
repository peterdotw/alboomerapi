package handlers

import (
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
	database.Dot.Exec(database.Db, "create-artists-table")
	database.Dot.Exec(database.Db, "create-albums-table")
	database.Dot.Exec(database.Db, "create-tracks-table")
}

func addAlbum(ID int, name, releaseDate, genre string) {
	database.Dot.Exec(database.Db, "create-album", name, ID, releaseDate, genre)
}

func init() {
	initDatabases()
	log.Println("Tests started")
}
