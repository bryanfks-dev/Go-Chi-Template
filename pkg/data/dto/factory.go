package basedto

type HTTPResponse[T any] struct {
	Success bool `json:"success"        example:"true"`
	Data    T    `json:"data,omitempty"`
}

type HTTPWithPaginationResponse[T any] struct {
	Success bool     `json:"success"        example:"true"`
	Meta    *MetaDto `json:"meta,omitempty"`
	Data    []T      `json:"data,omitempty"`
}

type ErrorHTTPResponse struct {
	Success bool      `json:"success" example:"false"`
	Error   *ErrorDto `json:"error"`
}

type ValidationErrorHTTPResponse struct {
	Success bool                `json:"success" example:"false"`
	Error   *ValidationErrorDto `json:"error"`
}

func NewHTTPResponse(data any) *HTTPResponse[any] {
	return &HTTPResponse[any]{
		Success: true,
		Data:    data,
	}
}

func NewHTTPWithPaginationResponse[T any](
	meta *MetaDto,
	data []T,
) *HTTPWithPaginationResponse[T] {
	return &HTTPWithPaginationResponse[T]{
		Success: true,
		Meta:    meta,
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
				Message: "validation_error",
			},
			Detail: detail,
		},
	}
}
