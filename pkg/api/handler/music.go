package handler

import (
	"encoding/json"
	"github.com/dafuqqqyunglean/music_library/pkg/models"
	"github.com/dafuqqqyunglean/music_library/pkg/service/music"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// CreateSong godoc
// @Summary Create song
// @Security ApiKeyAuth
// @Tags songs
// @Description create song
// @ID create-song
// @Accept  json
// @Produce  json
// @Param input body models.Song true "song info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Failure default {object} utility.ErrorResponse
// @Router /api/songs [post]
func CreateSong(ctx utility.AppContext, service music.MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		var input models.Song
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		id, err := service.Create(ctx, userId, input)
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"id": id,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

type getAllSongsResponse struct {
	Data []models.Song `json:"data"`
}

// GetAllSongs godoc
// @Summary Get All Songs
// @Security ApiKeyAuth
// @Tags songs
// @Description get all songs
// @ID get-all-songs
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllSongsResponse
// @Failure 400,404 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Failure default {object} utility.ErrorResponse
// @Router /api/songs [get]
func GetAllSongs(ctx utility.AppContext, service music.MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		songs, err := service.GetAll(ctx, userId)
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := getAllSongsResponse{
			Data: songs,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

// GetSongById godoc
// @Summary Get Song By id
// @Security ApiKeyAuth
// @Tags songs
// @Description get song by id
// @ID get-song-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Song
// @Failure 400,404 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Failure default {object} utility.ErrorResponse
// @Router /api/songs/:id [get]
func GetSongById(ctx utility.AppContext, service music.MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		song, err := service.GetById(ctx, userId, id)
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(song); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func DeleteSong(ctx utility.AppContext, service music.MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		err = service.Delete(ctx, userId, id)
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(utility.StatusResponse{Status: "ok"}); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func UpdateSong(ctx utility.AppContext, service music.MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserId").(int)

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		var input models.UpdateSongInput
		if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err = service.Update(ctx, userId, id, input); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(utility.StatusResponse{Status: "ok"}); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
