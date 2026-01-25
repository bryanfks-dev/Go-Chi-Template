package userdelivery

import (
	"net/http"
	basedto "skeleton/pkg/data/dto"
	"skeleton/pkg/utils"

	_ "skeleton/docs"
)

// @Summary List User Endpoint
// @Description Endpoint to list example items
// @Tags API / User
// @Produce json
// @Router /api/user [get]
// @Success 200 {object} basedto.SuccessHTTPResponse[any]
// @Failure 500 {object} basedto.ErrorHTTPResponse
func (h *UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	// Implementation here

	utils.WriteSuccessJSONResponse(
		w,
		http.StatusOK,
		basedto.NewSuccessHTTPResponse(nil, nil),
	)
}
