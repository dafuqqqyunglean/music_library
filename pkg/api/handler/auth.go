package handler

import (
	"encoding/json"
	"github.com/dafuqqqyunglean/music_library/pkg/models"
	"github.com/dafuqqqyunglean/music_library/pkg/service/auth"
	"github.com/dafuqqqyunglean/music_library/pkg/utility"
	"net/http"
)

// SignUp godoc
// @Summary Register a new user
// @Tags auth
// @Description Create a new user account
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body models.User true "User information"
// @Success 200 {object} map[string]interface{} "User ID"
// @Failure 400 {object} utility.ErrorResponse "Invalid input data"
// @Failure 500 {object} utility.ErrorResponse "Internal server error"
// @Router /auth/sign-up [post]
func SignUp(service auth.AuthorizationService, ctx utility.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input models.User

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		id, err := service.CreateUser(input)
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

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignIn godoc
// @Summary Log in an existing user
// @Tags auth
// @Description Authenticate a user and return a JWT token
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body signInInput true "User credentials"
// @Success 200 {object} map[string]interface{} "JWT token"
// @Failure 400 {object} utility.ErrorResponse "Invalid input data"
// @Failure 500 {object} utility.ErrorResponse "Internal server error"
// @Router /auth/sign-in [post]
func SignIn(service auth.AuthorizationService, ctx utility.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input signInInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusBadRequest, err.Error())
			return
		}

		token, err := service.GenerateToken(input.Username, input.Password)
		if err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"token": token,
		}
		if err = json.NewEncoder(w).Encode(response); err != nil {
			utility.NewErrorResponse(w, ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
