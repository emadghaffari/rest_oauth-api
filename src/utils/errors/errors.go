package errors

import "net/http"

// ResError struct
type ResError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// HandlerBagRequest func for ResError struct
func HandlerBagRequest(message string) *ResError {
	return &ResError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// HandlerNotFoundError func for ResError struct
func HandlerNotFoundError(message string) *ResError {
	return &ResError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// HandlerInternalServerError func for ResError struct
func HandlerInternalServerError(message string) *ResError {
	return &ResError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
