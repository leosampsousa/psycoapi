package errors

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	var error = Error{Message: message, Code: code}
	return &error
}