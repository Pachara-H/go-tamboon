package errors

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
	return e.data.Message
}

type unauthorizedError struct {
	data commonError
}

// Error for implement error interface
func (e *unauthorizedError) Error() string {
	return e.data.Message
}

type notFoundError struct {
	data commonError
}

// Error for implement error interface
func (e *notFoundError) Error() string {
	return e.data.Message
}

type unsupportedMediaTypeError struct {
	data commonError
}

// Error for implement error interface
func (e *unsupportedMediaTypeError) Error() string {
	return e.data.Message
}

type tooManyRequestsError struct {
	data commonError
}

// Error for implement error interface
func (e *tooManyRequestsError) Error() string {
	return e.data.Message
}

type internalServerError struct {
	data commonError
}

// Error for implement error interface
func (e *internalServerError) Error() string {
	return e.data.Message
}
