SELECT sg.id, sg.group_name, sg.song, sg.genre, sg.date, sg.lyrics, sg.link
FROM songs sg
INNER JOIN users_songs us on sg.id = us.song_id
WHERE us.user_id = $1 AND us.song_id = $2