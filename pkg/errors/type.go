package errors

import "fmt"

type commonError struct {
	Message string
	Status  int
	Code    int
}

type badRequestError struct {
	data commonError
}

// Error for implement error interface
func (e *badRequestError) Error() string {
	return fmt.Sprintf("%s (%d)", e.data.Message, e.data.Code)
}

type unauthorizedError struct {
	data commonError
}

// Error for implement error interface
func (e *unauthorizedError) Error() string {
	return fmt.Sprintf("%s (%d)", e.data.Message, e.data.Code)
}

type notFoundError struct {
	data commonError
}

// Error for implement error interface
func (e *notFoundError) Error() string {
	return fmt.Sprintf("%s (%d)", e.data.Message, e.data.Code)
}

type unsupportedMediaTypeError struct {
	data commonError
}

// Error for implement error interface
func (e *unsupportedMediaTypeError) Error() string {
	return fmt.Sprintf("%s (%d)", e.data.Message, e.data.Code)
}

type tooManyRequestsError struct {
	data commonError
}

// Error for implement error interface
func (e *tooManyRequestsError) Error() string {
	return fmt.Sprintf("%s (%d)", e.data.Message, e.data.Code)
}

type internalServerError struct {
	data commonError
}

// Error for implement error interface
func (e *internalServerError) Error() string {
	return fmt.Sprintf("%s (%d)", e.data.Message, e.data.Code)
}
