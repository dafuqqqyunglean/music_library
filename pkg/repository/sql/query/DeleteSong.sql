DELETE FROM songs sg USING users_songs us
WHERE sg.id = us.song_id AND us.user_id = $1 AND us.song_id = $2