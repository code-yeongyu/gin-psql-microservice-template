package custom_errors

// CustomError is a type for containing custom errors for http responses
type CustomError string

const (
	UnknownError = CustomError("UnknownError")
)
