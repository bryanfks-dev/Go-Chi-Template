package apperror

type AppError struct {
	Code int
	Err  error
}

func NewAppError(code int, err error) *AppError {
	return &AppError{
		Code: code,
		Err:  err,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}
