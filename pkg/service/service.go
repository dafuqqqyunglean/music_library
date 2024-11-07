package service

import (
	"github.com/dafuqqqyunglean/music_library/pkg/repository/sql"
	"github.com/dafuqqqyunglean/music_library/pkg/service/auth"
	"github.com/dafuqqqyunglean/music_library/pkg/service/music"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	AuthService  auth.AuthorizationService
	MusicService music.MusicService
}

func NewService(db *sqlx.DB, ctx utility.AppContext) *Service {
	authService := auth.NewAuthorizationService(sql.NewAuthorizationPostgres(db), ctx)
	musicService := music.NewMusicService(sql.NewMusicPostgres(db))
	return &Service{
		AuthService:  authService,
		MusicService: musicService,
	}
}
