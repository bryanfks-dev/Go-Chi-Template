package utils

import (
	"encoding/json"
	"net/http"
	basedto "skeleton/pkg/data/dto"
)

func WriteJSONResponse(
	w http.ResponseWriter,
	statusCode int,
	response any,
) {
	if response == nil {
		panic("response cannot be nil")
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func WriteErrorJSONResponse(
	w http.ResponseWriter,
	statusCode int,
	err error,
) {
	if statusCode < 400 || statusCode > 599 {
		panic("statusCode must be between 400 and 599")
	}

	if err == nil {
		panic("err cannot be nil")
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	response := basedto.NewErrorHTTPResponse(err.Error())
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
