package utility

import (
	"encoding/json"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, ctx AppContext, statusCode int, message string) {
	ctx.Logger.Errorf(message)

	errResponse := ErrorResponse{Message: message}

	jsonErrResponse, err := json.Marshal(errResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(statusCode)
	w.Write(jsonErrResponse)
}
