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
// @Router /public/health [get]
// @Success 200 {object} basedto.SuccessHTTPResponse[any]
// @Failure 500 {object} basedto.ErrorHTTPResponse
func (eh *EtcHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteSuccessJSONResponse(
		w,
		http.StatusOK,
		basedto.NewSuccessHTTPResponse(nil, nil),
	)
}

func (eh *EtcHandler) ErrorTest(w http.ResponseWriter, r *http.Request) {
}
