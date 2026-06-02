package errordelivery

import (
	"net/http"
	apperror "skeleton/pkg/error"
	serverutils "skeleton/pkg/server/utils"
)

func (h *ErrorHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	err := apperror.NewAppError(
		http.StatusNotFound,
		apperror.ErrRouteNotFound,
	)
	serverutils.WriteErrorJSONResponse(w, http.StatusNotFound, err)
}

func (h *ErrorHandler) MethodNotAllowed(
	w http.ResponseWriter,
	r *http.Request,
) {
	err := apperror.NewAppError(
		http.StatusMethodNotAllowed,
		apperror.ErrMethodNotAllowed,
	)
	serverutils.WriteErrorJSONResponse(w, http.StatusMethodNotAllowed, err)
}
