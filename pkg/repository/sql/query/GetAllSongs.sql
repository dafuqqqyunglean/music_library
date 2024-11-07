SELECT sg.id, sg.group_name, sg.song, sg.genre, sg.date, sg.lyrics, sg.link
FROM songs sg
INNER JOIN users_songs us ON us.song_id = sg.id
WHERE us.user_id = $1