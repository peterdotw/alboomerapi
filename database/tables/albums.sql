-- name: create-albums-table
CREATE TABLE IF NOT EXISTS albums (
  album_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  album_name VARCHAR(128) NOT NULL,
  artist_id INT UNSIGNED NOT NULL,
  release_date DATE NOT NULL,
  genre VARCHAR(64) NOT NULL,
  PRIMARY KEY (album_id),
  FOREIGN KEY (artist_id) REFERENCES artists(artist_id) ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8 DEFAULT COLLATE = utf8_unicode_ci;

-- name: create-album
INSERT INTO albums VALUES (default, ?, ?, ?, ?);

-- name: select-albums
SELECT al.album_id, al.album_name, ar.artist_name, al.release_date, al.genre FROM albums AS al
JOIN artists AS ar ON (ar.artist_id = al.artist_id);

-- name: select-album
SELECT al.album_id, al.album_name, ar.artist_name, al.release_date, al.genre FROM albums AS al
JOIN artists AS ar ON (ar.artist_id = al.artist_id) WHERE album_id = ?;

-- name: update-album
UPDATE albums SET album_name = ?, artist_id = ?, release_date = ?, genre = ? WHERE album_id = ?;

-- name: delete-album
DELETE FROM albums WHERE album_id = ?;