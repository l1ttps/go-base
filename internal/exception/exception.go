package exception

type HttpExceptionResponse struct {
	Status  int
	Message string
}

// HttpException creates a new HttpExceptionResponse with the given status code and message.
//
// Parameters:
// - statusCode: the HTTP status code.
// - message: the error message.
//
// Return:
// - HttpExceptionResponse: the generated HttpExceptionResponse.
func HttpException(status int, message string) HttpExceptionResponse {
	return HttpExceptionResponse{
		Status:  status,
		Message: message,
	}
}
