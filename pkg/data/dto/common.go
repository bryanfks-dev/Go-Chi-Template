package basedto

type ValidationErrorDTO map[string]string
type PaginationResponseDTO struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	SizePerPage int `json:"size_per_page"`
}

type MetaResponseDTO struct {
	*PaginationResponseDTO
}

type ErrorResponseDTO struct {
	Message string `json:"message"`
}

type ValidationErrorResponseDTO struct {
	ErrorResponseDTO
	Detail ValidationErrorDTO `json:"detail"`
}
