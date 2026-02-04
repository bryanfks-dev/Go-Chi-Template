package utils

import (
	"encoding/json"
	"net/http"
	basedto "skeleton/pkg/data/dto"
	apperror "skeleton/pkg/error"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ReadJSONRequest[T any](
	r *http.Request,
	dest T,
) error {
	if r.Body == nil {
		return apperror.NewAppError(
			http.StatusBadRequest,
			apperror.ErrMissingRequestBody,
		)
	}

	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return apperror.NewAppError(
			http.StatusBadRequest,
			apperror.ErrInvalidRequestBody,
		)
	}
	defer r.Body.Close()

	return nil
}

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
		WriteErrorJSONResponse(w, err)
	}
}

func WriteNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func WriteErrorJSONResponse(
	w http.ResponseWriter,
	err error,
) {
	if err == nil {
		panic("err cannot be nil")
	}

	appErr, ok := err.(*apperror.AppError)
	if !ok {
		appErr = apperror.NewAppError(
			http.StatusInternalServerError,
			apperror.ErrInternalServer,
		)
	}

	if appErr.Code < 400 || appErr.Code > 599 {
		panic("statusCode must be between 400 and 599")
	}

	w.WriteHeader(appErr.Code)
	w.Header().Set("Content-Type", "application/json")

	res := basedto.NewErrorHTTPResponse(appErr.Error())
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteValidationErrorJSONResponse(
	w http.ResponseWriter,
	err error,
) {
	if err == nil {
		panic("err cannot be nil")
	}

	validationErr := map[string]string{}
	if errs, ok := err.(validation.Errors); ok {
		for field, fieldErr := range errs {
			if fieldErr == nil {
				continue
			}
			validationErr[field] = fieldErr.Error()
		}
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Header().Set("Content-Type", "application/json")

	res := basedto.NewValidationErrorHTTPResponse(validationErr)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		WriteErrorJSONResponse(w, err)
	}
}
