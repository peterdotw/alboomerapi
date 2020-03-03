-- name: create-tracks-table
CREATE TABLE IF NOT EXISTS tracks (
  track_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  track_name VARCHAR(128) NOT NULL,
  album_id INT UNSIGNED NOT NULL,
  PRIMARY KEY (track_id),
  FOREIGN KEY (album_id) REFERENCES albums(album_id) ON DELETE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8 DEFAULT COLLATE = utf8_unicode_ci;

-- name: create-track
INSERT INTO tracks VALUES (default, ?, ?);

-- name: select-tracks
SELECT * FROM tracks;

-- name: select-track
SELECT * FROM tracks WHERE track_id = ?;

-- name: update-track
UPDATE tracks SET track_name = ?, album_id = ? WHERE track_id = ?;

-- name: delete-track
DELETE FROM tracks WHERE track_id = ?;