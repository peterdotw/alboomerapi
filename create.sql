CREATE TABLE albums (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  artist_name VARCHAR(128) NOT NULL,
  album_name VARCHAR(128) NOT NULL,
  genre VARCHAR(64) NOT NULL,
  release_date DATE NOT NULL
) engine = innodb DEFAULT charset = utf8 DEFAULT COLLATE = utf8_unicode_ci;

CREATE TABLE tracks (
  track VARCHAR(128) NOT NULL,
  artist_name VARCHAR(128) NOT NULL
)

CREATE TABLE genres (
  genre VARCHAR(64) NOT NULL
)