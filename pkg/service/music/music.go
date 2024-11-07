package music

import (
	"github.com/dafuqqqyunglean/music_library/pkg/models"
	"github.com/dafuqqqyunglean/music_library/pkg/repository/sql"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
)

type MusicService interface {
	Create(ctx utility.AppContext, userId int, song models.Song) (int, error)
	GetAll(ctx utility.AppContext, userId int) ([]models.Song, error)
	GetById(ctx utility.AppContext, userId, songId int) (models.Song, error)
	Update(ctx utility.AppContext, userId, songId int, input models.UpdateSongInput) error
	Delete(ctx utility.AppContext, userId, songId int) error
}

type ImplMusic struct {
	repo sql.MusicRepository
}

func NewMusicService(repo sql.MusicRepository) *ImplMusic {
	return &ImplMusic{repo: repo}
}

func (s *ImplMusic) Create(ctx utility.AppContext, userId int, song models.Song) (int, error) {
	return s.repo.Create(ctx, userId, song)
}

func (s *ImplMusic) GetAll(ctx utility.AppContext, userId int) ([]models.Song, error) {
	items, err := s.repo.GetAll(ctx, userId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *ImplMusic) GetById(ctx utility.AppContext, userId, songId int) (models.Song, error) {
	song, err := s.repo.GetById(ctx, userId, songId)
	if err != nil {
		return song, err
	}
	return song, nil
}

func (s *ImplMusic) Update(ctx utility.AppContext, userId, songId int, input models.UpdateSongInput) error {
	err := input.Validate()
	if err != nil {
		return err
	}

	err = s.repo.Update(ctx, userId, songId, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *ImplMusic) Delete(ctx utility.AppContext, userId, songId int) error {
	err := s.repo.Delete(ctx, userId, songId)
	if err != nil {
		return err
	}

	return nil
}
