package etcdelivery

import (
	"net/http"
	basedto "skeleton/pkg/data/dto"
	"skeleton/pkg/utils"

	_ "skeleton/docs"
)

// @Summary Health Check
// @Description Endpoint to check the health of the service
// @Tags Public API / Etc
// @Produce json
// @Router /public/api/health [get]
// @Success 200 {object} basedto.HTTPResponse[any]
// @Failure 500 {object} basedto.ErrorHTTPResponse
func (h *EtcHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSONResponse(
		w,
		http.StatusOK,
		basedto.NewHTTPResponse(nil),
	)
}
