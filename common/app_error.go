package common

import "errors"

type AppError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	RootError  error  `json:"root_error"`
}

func (a *AppError) Error() string {
	return a.GetRootError().Error()
}

func (a *AppError) GetRootError() error {
	if err, ok := a.RootError.(*AppError); ok {
		return err.GetRootError()
	}
	return a.RootError
}

func NewErrorResponse(statusCode int, message string, rootError error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		RootError:  rootError,
	}
}

func LoginErrorResponse(statusCode int, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		RootError:  errors.New("Wrong email or password"),
	}
}
