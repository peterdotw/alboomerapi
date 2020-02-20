CREATE TABLE albums (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  album_name VARCHAR(128) NOT NULL,
  artist_name VARCHAR(128) NOT NULL,
  release_date DATE NOT NULL,
  genre VARCHAR(64) NOT NULL
) engine = innodb DEFAULT charset = utf8 DEFAULT COLLATE = utf8_unicode_ci;