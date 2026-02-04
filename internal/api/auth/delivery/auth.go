package authdelivery

import (
	"net/http"
	authdto "skeleton/internal/api/auth/data/dto"
	basedto "skeleton/pkg/data/dto"
	"skeleton/pkg/utils"

	_ "skeleton/docs"
)

// @Summary Login Endpoint
// @Description Endpoint to login user
// @Tags API / Auth
// @Accept json
// @Produce json
// @Router /api/auth/login [post]
// @Param request body authdto.PostAuthLoginRequestDTO true "Request Body"
// @Success 200 {object} basedto.HTTPResponse[authdto.PostAuthLoginResponseDTO]
// @Failure 400 {object} basedto.ErrorHTTPResponse "Bad Request"
// @Failure 401 {object} basedto.ErrorHTTPResponse "Unauthorized"
// @Failure 422 {object} basedto.ValidationErrorHTTPResponse "Unprocessable Entity"
// @Failure 500 {object} basedto.ErrorHTTPResponse "Internal Server Error"
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req authdto.PostAuthLoginRequestDTO
	if err := utils.ReadJSONRequest(r, &req); err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}
	if err := req.Validate(); err != nil {
		utils.WriteValidationErrorJSONResponse(w, err)
		return
	}

	userAgent := r.UserAgent()
	user, refreshToken, accessToken, csrfToken, xsrfToken, err :=
		h.authUc.ProcessUserLogin(
			r.Context(),
			req.Email,
			req.Password,
			userAgent,
		)
	if err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	resData := authdto.NewPostAuthLoginResponseDTO(
		user,
		accessToken,
		refreshToken,
		csrfToken,
		xsrfToken,
	)
	res := basedto.NewHTTPResponse(resData)
	utils.WriteJSONResponse(w, http.StatusOK, res)
}

// @Summary Logout Endpoint
// @Description Endpoint to logout user
// @Tags API / Auth
// @Accept json
// @Produce json
// @Router /api/auth/logout [post]
// @Param request body authdto.PostAuthLogoutRequestDTO true "Request Body"
// @Success 200 {object} basedto.HTTPResponse[any]
// @Failure 400 {object} basedto.ErrorHTTPResponse "Bad Request"
// @Failure 401 {object} basedto.ErrorHTTPResponse "Unauthorized"
// @Failure 422 {object} basedto.ValidationErrorHTTPResponse "Unprocessable Entity"
// @Failure 500 {object} basedto.ErrorHTTPResponse "Internal Server Error"
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req authdto.PostAuthLogoutRequestDTO
	if err := utils.ReadJSONRequest(r, &req); err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	if err := req.Validate(); err != nil {
		utils.WriteValidationErrorJSONResponse(w, err)
		return
	}

	err := h.authUc.ProcessUserLogout(r.Context(), req.RefreshToken)
	if err != nil {
		utils.WriteErrorJSONResponse(w, err)
		return
	}

	res := basedto.NewHTTPResponse(nil)
	utils.WriteJSONResponse(w, http.StatusOK, res)
}
