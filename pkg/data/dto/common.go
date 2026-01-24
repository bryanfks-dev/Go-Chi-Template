package basedto

type ValidationError map[string]string

type MetaDto struct {
	Total       int64 `json:"total,omitempty"`
	Page        int   `json:"page,omitempty"`
	SizePerPage int   `json:"size_per_page,omitempty"`
}

type ErrorDto struct {
	Message string `json:"message"`
}

type ValidationErrorDto struct {
	ErrorDto
	Detail ValidationError `json:"detail"`
}
