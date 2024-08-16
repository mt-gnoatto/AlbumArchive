DROP TABLE IF EXISTS album;

CREATE TABLE album (
    name VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    genre VARCHAR(100) NOT NULL,
    score INT DEFAULT 0 CHECK (score >= 0 AND score <= 5),
    liked BOOLEAN NOT NULL DEFAULT FALSE,
    played BOOLEAN NOT NULL DEFAULT FALSE
);

