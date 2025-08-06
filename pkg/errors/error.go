// Package errors is a common utility error type
package errors

import (
	"github.com/gofiber/fiber/v2"
)

const (
	msgBadRequestError      = "Invalid data or parameter"
	msgUnauthorizedError    = "Unauthorized"
	msgNotFoundError        = "Data was not found"
	msgTooManyRequestsError = "Too many requests"
	msgInternalServerError  = "Internal system error"
)

var getMessage = map[int]string{
	fiber.StatusBadRequest:          msgBadRequestError,
	fiber.StatusUnauthorized:        msgUnauthorizedError,
	fiber.StatusNotFound:            msgNotFoundError,
	fiber.StatusTooManyRequests:     msgTooManyRequestsError,
	fiber.StatusInternalServerError: msgInternalServerError,
}

// NewBadRequestError for initial new bad request error
func NewBadRequestError(code int, message ...string) error {
	return &badRequestError{
		data: setErrorData(fiber.StatusBadRequest, code, message...),
	}
}

// NewUnauthorizedError for initial new unauthorized error
func NewUnauthorizedError(code int, message ...string) error {
	return &unauthorizedError{
		data: setErrorData(fiber.StatusUnauthorized, code, message...),
	}
}

// NewNotFoundError for initial new not found error
func NewNotFoundError(code int, message ...string) error {
	return &notFoundError{
		data: setErrorData(fiber.StatusNotFound, code, message...),
	}
}

// NewTooManyRequestsError for initial new too many requests error
func NewTooManyRequestsError(code int, message ...string) error {
	return &tooManyRequestsError{
		data: setErrorData(fiber.StatusTooManyRequests, code, message...),
	}
}

// NewInternalServerError for initial new internal server error
func NewInternalServerError(code int, message ...string) error {
	return &internalServerError{
		data: setErrorData(fiber.StatusInternalServerError, code, message...),
	}
}

func setErrorData(status, code int, message ...string) commonError {
	msg := getMessage[status]
	if len(message) > 0 {
		msg = message[0]
	}
	if msg == "" {
		msg = "Something went wrong" // default message
	}

	return commonError{
		Message: msg,
		Status:  status,
		Code:    code,
	}
}
