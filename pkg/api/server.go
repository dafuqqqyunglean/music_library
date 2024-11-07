package api

import (
	"context"
	_ "github.com/dafuqqqyunglean/music_library/docs"
	"github.com/dafuqqqyunglean/music_library/pkg/api/handler"
	"github.com/dafuqqqyunglean/music_library/pkg/api/middlewares"
	"github.com/dafuqqqyunglean/music_library/pkg/service/auth"
	"github.com/dafuqqqyunglean/music_library/pkg/service/music"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"time"
)

const (
	maxHeaderBytes = 1 << 20 // 1 MB
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewServer(middleware *middlewares.UserAuthMiddleware) *Server {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.UserAuth)

	return &Server{
		httpServer: &http.Server{
			Addr:           ":8080",
			MaxHeaderBytes: maxHeaderBytes,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			Handler:        router,
		},
		router: router,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) HandleAuth(service auth.AuthorizationService, ctx utility.AppContext) {
	s.router.HandleFunc("/auth/sign-up/", handler.SignUp(service, ctx)).Methods(http.MethodPost)
	s.router.HandleFunc("/auth/sign-in/", handler.SignIn(service, ctx)).Methods(http.MethodPost)
}

func (s *Server) HandleMusic(service music.MusicService, ctx utility.AppContext) {
	s.router.HandleFunc("/songs/", handler.CreateSong(ctx, service)).Methods(http.MethodPost)
	s.router.HandleFunc("/songs/", handler.GetAllSongs(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/songs/{id}", handler.GetSongById(ctx, service)).Methods(http.MethodGet)
	s.router.HandleFunc("/songs/{id}", handler.DeleteSong(ctx, service)).Methods(http.MethodDelete)
	s.router.HandleFunc("/songs/{id}", handler.UpdateSong(ctx, service)).Methods(http.MethodPut)
}
