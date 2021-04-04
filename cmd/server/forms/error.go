package forms

type ErrorResponse struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
	Detail    string `json:"detail"`
}
