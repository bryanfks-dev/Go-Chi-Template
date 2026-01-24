package basedto

type SuccessHTTPResponse[T any] struct {
	Success bool     `json:"success" example:"true"`
	Meta    *MetaDto `json:"meta,omitempty"`
	Data    T        `json:"data,omitempty"`
}

type ErrorHTTPResponse struct {
	Success bool      `json:"success" example:"false"`
	Error   *ErrorDto `json:"error"`
}

type ValidationErrorHTTPResponse struct {
	Success bool                `json:"success" example:"false"`
	Error   *ValidationErrorDto `json:"error"`
}

func NewSuccessHTTPResponse(
	data any,
	meta *MetaDto,
) *SuccessHTTPResponse[any] {
	return &SuccessHTTPResponse[any]{
		Meta:    meta,
		Success: true,
		Data:    data,
	}
}

func NewErrorHTTPResponse(message string) *ErrorHTTPResponse {
	return &ErrorHTTPResponse{
		Success: false,
		Error: &ErrorDto{
			Message: message,
		},
	}
}

func NewValidationErrorHTTPResponse(
	detail ValidationError,
) *ValidationErrorHTTPResponse {
	return &ValidationErrorHTTPResponse{
		Success: false,
		Error: &ValidationErrorDto{
			ErrorDto: ErrorDto{
				Message: "validation error",
			},
			Detail: detail,
		},
	}
}
