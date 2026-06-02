package basedto

type HTTPResponse[T any] struct {
	Success bool `json:"success" example:"true"`
	Data    T    `json:"data"`
}

type HTTPWithPaginationResponse[T any] struct {
	Success bool             `json:"success"        example:"true"`
	Meta    *MetaResponseDTO `json:"meta,omitempty"`
	Data    []T              `json:"data"`
}

type ErrorHTTPResponse struct {
	Success bool              `json:"success" example:"false"`
	Error   *ErrorResponseDTO `json:"error"`
}

type ValidationErrorHTTPResponse struct {
	Success bool                        `json:"success" example:"false"`
	Error   *ValidationErrorResponseDTO `json:"error"`
}

func NewHTTPResponse(data any) *HTTPResponse[any] {
	return &HTTPResponse[any]{
		Success: true,
		Data:    data,
	}
}

func NewHTTPWithPaginationResponse[T any](
	meta *MetaResponseDTO,
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
		Error: &ErrorResponseDTO{
			Message: message,
		},
	}
}

func NewValidationErrorHTTPResponse(
	detail ValidationErrorDTO,
) *ValidationErrorHTTPResponse {
	return &ValidationErrorHTTPResponse{
		Success: false,
		Error: &ValidationErrorResponseDTO{
			ErrorResponseDTO: ErrorResponseDTO{
				Message: "VALIDATION_ERROR",
			},
			Detail: detail,
		},
	}
}

func NewMetaResponseDTOFromPaginationResultDTO(
	pagination PaginationResultDTO,
) *MetaResponseDTO {
	paginationRes := NewPaginationResponseDTOFromPaginationResultDTO(pagination)
	return &MetaResponseDTO{
		paginationRes,
	}
}

func NewPaginationResponseDTOFromPaginationResultDTO(
	pagination PaginationResultDTO,
) *PaginationResponseDTO {
	return &PaginationResponseDTO{
		TotalItems:  pagination.TotalItems,
		TotalPages:  pagination.GetTotalPages(),
		SizePerPage: pagination.SizePerPage,
	}
}
