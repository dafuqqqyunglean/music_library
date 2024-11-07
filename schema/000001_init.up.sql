CREATE TABLE users
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    username      VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE songs
(
    id           SERIAL PRIMARY KEY,
    group_name   VARCHAR(255) NOT NULL,
    song         VARCHAR(255) NOT NULL,
    genre        VARCHAR(100),
    date         DATE,
    lyrics       TEXT,
    link         VARCHAR(255)
);

CREATE TABLE users_songs
(
    id       SERIAL PRIMARY KEY,
    user_id  INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    song_id  INT REFERENCES songs(id) ON DELETE CASCADE NOT NULL
);