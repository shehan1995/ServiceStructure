package error_handler

import "net/http"

const ServerError = `Internal Server Error`

// Error is the default error message that is generated.
// This is used for all the error cases.
type Error struct {
	// Code denotes http status code
	Code int `json:"code"`
	// Message denotes the error message in high-level
	Message string `json:"message"`
}

type ErrorHandler func(error) Error

func (e Error) Error() string {
	return e.Message
}

func ValidationErr(err error) Error {
	return createError(http.StatusUnprocessableEntity, err)
}

func RequestErr(err error) Error {
	return createError(http.StatusBadRequest, err)
}

func createError(code int, err error) Error {
	return Error{Code: code, Message: err.Error()}
}
