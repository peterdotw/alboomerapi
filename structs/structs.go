package structs

// Album : Structure for single album
type Album struct {
	ID          int    `json:"album_id"`
	Name        string `json:"album_name"`
	ArtistName  string `json:"artist_name"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
}

// Artist : Structure for single artist
type Artist struct {
	ID         int    `json:"artist_id"`
	ArtistName string `json:"artist_name"`
}

// Track : Structure for single track
type Track struct {
	ID        int    `json:"track_id"`
	TrackName string `json:"track_name"`
	AlbumID   int    `json:"album_id"`
}

// Albums : Structure for albums
type Albums struct {
	Albums []Album `json:"albums"`
}

// Artists : Structure for artists
type Artists struct {
	Artists []Artist `json:"artists"`
}

// Tracks : Structure for tracks
type Tracks struct {
	Tracks []Track `json:"tracks"`
}
