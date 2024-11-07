package sql

import (
	_ "embed"
	"fmt"
	"github.com/dafuqqqyunglean/music_library/pkg/models"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"github.com/jmoiron/sqlx"
	"strings"
)

type MusicRepository interface {
	Create(ctx utility.AppContext, userId int, song models.Song) (int, error)
	GetAll(ctx utility.AppContext, userId int) ([]models.Song, error)
	GetById(ctx utility.AppContext, userId, songId int) (models.Song, error)
	Update(ctx utility.AppContext, userId, songId int, input models.UpdateSongInput) error
	Delete(ctx utility.AppContext, userId, songId int) error
}

type MusicPostgres struct {
	db *sqlx.DB
}

func NewMusicPostgres(db *sqlx.DB) *MusicPostgres {
	return &MusicPostgres{db: db}
}

//go:embed query/CreateSong.sql
var createSong string

//go:embed query/CreateUsersSongs.sql
var createUsersSongs string

func (r *MusicPostgres) Create(ctx utility.AppContext, userId int, song models.Song) (int, error) {
	tx, err := r.db.BeginTx(ctx.Ctx, nil)
	if err != nil {
		return 0, err
	}

	var songId int
	row := tx.QueryRow(createSong, song.Group, song.Song, song.Genre, song.Date, song.Lyrics, song.Link)
	if err := row.Scan(&songId); err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(createUsersSongs, userId, songId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return songId, tx.Commit()
}

//go:embed query/GetAllSongs.sql
var getAllSongs string

func (r *MusicPostgres) GetAll(ctx utility.AppContext, userId int) ([]models.Song, error) {
	var songs []models.Song

	if err := r.db.SelectContext(ctx.Ctx, &songs, getAllSongs, userId); err != nil {
		return nil, err
	}

	return songs, nil
}

//go:embed query/GetSongById.sql
var getSongById string

func (r *MusicPostgres) GetById(ctx utility.AppContext, userId, songId int) (models.Song, error) {
	var song models.Song

	err := r.db.GetContext(ctx.Ctx, &song, getSongById, userId, songId)

	return song, err
}

//go:embed query/UpdateSong.sql
var updateSong string

func (r *MusicPostgres) Update(ctx utility.AppContext, userId, songId int, input models.UpdateSongInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Group != nil {
		setValues = append(setValues, fmt.Sprintf("group_name = $%d", argId))
		args = append(args, *input.Group)
		argId++
	}

	if input.Song != nil {
		setValues = append(setValues, fmt.Sprintf("song = $%d", argId))
		args = append(args, *input.Song)
		argId++
	}

	if input.Genre != nil {
		setValues = append(setValues, fmt.Sprintf("genre = $%d", argId))
		args = append(args, *input.Genre)
		argId++
	}

	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date = $%d", argId))
		args = append(args, *input.Date)
		argId++
	}

	if input.Lyrics != nil {
		setValues = append(setValues, fmt.Sprintf("lyrics = $%d", argId))
		args = append(args, *input.Lyrics)
		argId++
	}

	if input.Link != nil {
		setValues = append(setValues, fmt.Sprintf("link = $%d", argId))
		args = append(args, *input.Link)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(updateSong, setQuery, argId, argId+1)
	args = append(args, songId, userId)

	_, err := r.db.ExecContext(ctx.Ctx, query, args...)
	return err
}

//go:embed query/DeleteSong.sql
var deleteSong string

func (r *MusicPostgres) Delete(ctx utility.AppContext, userId, songId int) error {
	_, err := r.db.ExecContext(ctx.Ctx, deleteSong, userId, songId)

	return err
}
