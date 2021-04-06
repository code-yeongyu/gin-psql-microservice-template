package errors

// CustomError is a type for containing custom errors for http responses
type CustomError string

const (
	ValidationError     = CustomError("ValidationError")
	DuplicateEmailError = CustomError("DuplicateEmailError")
)
