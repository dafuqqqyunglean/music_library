INSERT INTO songs (group_name, song, genre, date, lyrics, link) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id