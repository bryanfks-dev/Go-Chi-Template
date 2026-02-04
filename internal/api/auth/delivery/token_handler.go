package authdelivery

import (
	"net/http"
	authdto "skeleton/internal/api/auth/data/dto"
	basedto "skeleton/pkg/data/dto"
	"skeleton/pkg/utils"

	_ "skeleton/docs"
)

// @Summary Refresh Access Token Endpoint
// @Description Endpoint to refresh access token
// @Tags API / Auth
// @Accept json
// @Produce json
// @Router /api/auth/token/refresh [post]
// @Param request body authdto.PostAuthTokenRefreshRequestDTO true "Request Body"
// @Success 200 {object} basedto.HTTPResponse[authdto.PostAuthTokenRefreshResponseDTO]
// @Failure 400 {object} basedto.ErrorHTTPResponse "Bad Request"
// @Failure 401 {object} basedto.ErrorHTTPResponse "Unauthorized"
// @Failure 500 {object} basedto.ErrorHTTPResponse "Internal Server Error"
func (h *AuthHandler) RefreshAccessToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req authdto.PostAuthTokenRefreshRequestDTO
	if err := utils.ReadJSONRequest(r, &req); err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	refreshClaims, err := h.authUc.DecodeRefreshToken(
		r.Context(),
		req.RefreshToken,
	)
	if err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	accessToken, err := h.authUc.GenerateAccessToken(
		r.Context(),
		refreshClaims,
	)
	if err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	resData := authdto.NewPostAuthTokenRefreshResponseDTO(accessToken)
	res := basedto.NewHTTPResponse(resData)
	utils.WriteJSONResponse(w, http.StatusOK, res)
}
