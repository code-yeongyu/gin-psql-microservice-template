package errors

// Messages contains the error message for each error type
var Messages = map[CustomError]string{
	ValidationError:     "Validation error.",
	DuplicateEmailError: "Interest with the given email already exists.",
}
