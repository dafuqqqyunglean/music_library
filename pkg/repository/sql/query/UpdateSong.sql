UPDATE songs sg
SET %s
FROM users_songs us
WHERE sg.id = us.song_id AND us.song_id = $%d AND us.user_id = $%d