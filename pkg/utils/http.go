package utils

import (
	"encoding/json"
	"net/http"
	basedto "skeleton/pkg/data/dto"
)

func WriteSuccessJSONResponse(
	w http.ResponseWriter,
	statusCode int,
	response *basedto.SuccessHTTPResponse[any],
) {
	if statusCode < 200 || statusCode >= 300 {
		panic("statusCode must be a success code (2xx)")
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if response == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteErrorJSONResponse(
	w http.ResponseWriter,
	statusCode int,
	response *basedto.ErrorHTTPResponse,
) {
	if statusCode < 400 || statusCode >= 600 {
		panic("statusCode must be an error code (4xx or 5xx)")
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if response == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
