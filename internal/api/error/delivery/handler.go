package errordelivery

import (
	"net/http"
	apperror "skeleton/pkg/error"
	"skeleton/pkg/utils"
)

func (h *ErrorHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	utils.WriteErrorJSONResponse(
		w,
		apperror.NewAppError(
			http.StatusNotFound,
			apperror.ErrRouteNotFound,
		),
	)
}

func (h *ErrorHandler) MethodNotAllowed(
	w http.ResponseWriter,
	r *http.Request,
) {
	utils.WriteErrorJSONResponse(
		w,
		apperror.NewAppError(
			http.StatusMethodNotAllowed,
			apperror.ErrMethodNotAllowed,
		),
	)
}
