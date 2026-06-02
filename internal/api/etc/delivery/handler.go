package etcdelivery

import (
	"net/http"
	basedto "skeleton/pkg/data/dto"
	serverutils "skeleton/pkg/server/utils"

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
	response := basedto.NewHTTPResponse(nil)
	serverutils.WriteJSONResponse(w, http.StatusOK, response)
}
