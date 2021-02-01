package errorhandling

import "strconv"

type InternalError struct {
	Code    int
	Message string
}

func NewInternalError(code int) InternalError {
	return InternalError{
		code,
		ErrorDescription(code),
	}
}

func NewInternalErrorWithCustomizedMessage(code int, message string) InternalError {
	return InternalError{
		code,
		message,
	}
}

// Error is the implementation to get the error message
func (ie InternalError) Error() string {
	return strconv.Itoa(ie.Code) + " " + ie.Message
}
