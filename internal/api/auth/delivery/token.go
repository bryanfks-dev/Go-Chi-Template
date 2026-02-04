package authdelivery

import (
	"net/http"
	authdto "skeleton/internal/api/auth/data/dto"
	basedto "skeleton/pkg/data/dto"
	"skeleton/pkg/utils"

	_ "skeleton/docs"
)

// @Summary Refresh Authentication Tokens Endpoint
// @Description Endpoint to refresh Access, CSRF, and XSRF tokens
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

	accessToken, csrfToken, xsrfToken, err := h.authUc.RefreshAuthToken(
		r.Context(),
		req.RefreshToken,
	)
	if err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	resData := authdto.NewPostAuthTokenRefreshResponseDTO(
		accessToken,
		csrfToken,
		xsrfToken,
	)
	res := basedto.NewHTTPResponse(resData)
	utils.WriteJSONResponse(w, http.StatusOK, res)
}
