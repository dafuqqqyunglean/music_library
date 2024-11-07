package main

import (
	"context"
	"github.com/dafuqqqyunglean/music_library/config"
	"github.com/dafuqqqyunglean/music_library/pkg/api"
	"github.com/dafuqqqyunglean/music_library/pkg/api/middlewares"
	"github.com/dafuqqqyunglean/music_library/pkg/repository"
	"github.com/dafuqqqyunglean/music_library/pkg/service"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type App struct {
	ctx      utility.AppContext
	srv      *api.Server
	cfg      config.Config
	postgres *sqlx.DB
}

func NewApp(context context.Context, logger *zap.SugaredLogger, cfg config.Config) *App {
	return &App{
		ctx:      *utility.NewAppContext(context, logger),
		cfg:      cfg,
		postgres: repository.NewPostgresDB(cfg.Postgres, logger),
	}
}

func (a *App) InitService() {
	newService := service.NewService(a.postgres, a.ctx)
	a.srv = api.NewServer(middlewares.NewUserAuthMiddleware(newService.AuthService, a.ctx))
	a.srv.HandleAuth(newService.AuthService, a.ctx)
	a.srv.HandleMusic(newService.MusicService, a.ctx)
}

func (a *App) Run() error {
	go func() {
		err := a.srv.Run()
		if err != nil {
			a.ctx.Logger.Fatalf("error running http server: %s", err.Error())
		}
	}()

	a.ctx.Logger.Info("server running")
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	err := a.srv.Shutdown(ctx)
	if err != nil {
		a.ctx.Logger.Errorf("failed to disconnect from server %v", err)
		return err
	}

	err = a.postgres.Close()
	if err != nil {
		a.ctx.Logger.Errorf("failed to disconnect from bd %v", err)
	}

	a.ctx.Logger.Info("server shutdown successfully")
	return nil
}
