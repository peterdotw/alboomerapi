-- name: create-artists-table
CREATE TABLE IF NOT EXISTS artists (
  artist_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  artist_name VARCHAR(128) NOT NULL,
  PRIMARY KEY (artist_id),
  UNIQUE KEY (artist_name)
) ENGINE = INNODB DEFAULT CHARSET = utf8 DEFAULT COLLATE = utf8_unicode_ci;

-- name: create-artist
INSERT INTO artists VALUES (default, ?);

-- name: select-artists
SELECT * FROM artists ORDER BY artist_id ASC;

-- name: select-artist-id
SELECT artist_id FROM artists WHERE artist_name = ?;

-- name: update-artist
UPDATE artists SET artist_name = ? WHERE artist_id = ?;

-- name: delete-artist
DELETE FROM artists WHERE artist_id = ?;