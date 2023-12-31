package exception

import "net/http"

type HttpExceptionResponse struct {
	Status  int
	Message string
}

// HttpException creates a new HttpExceptionResponse with the given status and message.
//
// The status parameter specifies the HTTP status code for the exception.
//
// The message parameter specifies the optional custom message for the exception. If no
// message is provided, a default message will be used based on the status code.
//
// Returns an HttpExceptionResponse object representing the exception.
func HttpException(status int, message string) HttpExceptionResponse {
	if message == "" {
		message = http.StatusText(status)
	}
	return HttpExceptionResponse{
		Status:  status,
		Message: message,
	}
}

// Generated by Chat GPT :)

func CreateCustomException(status int, message ...string) HttpExceptionResponse {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	return HttpException(status, msg)
}

func BadRequestException(message ...string) HttpExceptionResponse {
	return CreateCustomException(400, message...)
}

func UnauthorizedException(message ...string) HttpExceptionResponse {
	return CreateCustomException(401, message...)
}

func NotFoundException(message ...string) HttpExceptionResponse {
	return CreateCustomException(404, message...)
}

func ForbiddenException(message ...string) HttpExceptionResponse {
	return CreateCustomException(403, message...)
}

func NotAcceptableException(message ...string) HttpExceptionResponse {
	return CreateCustomException(406, message...)
}

func RequestTimeoutException(message ...string) HttpExceptionResponse {
	return CreateCustomException(408, message...)
}

func ConflictException(message ...string) HttpExceptionResponse {
	return CreateCustomException(409, message...)
}

func GoneException(message ...string) HttpExceptionResponse {
	return CreateCustomException(410, message...)
}

func HttpVersionNotSupportedException(message ...string) HttpExceptionResponse {
	return CreateCustomException(505, message...)
}

func PayloadTooLargeException(message ...string) HttpExceptionResponse {
	return CreateCustomException(413, message...)
}

func UnsupportedMediaTypeException(message ...string) HttpExceptionResponse {
	return CreateCustomException(415, message...)
}

func UnprocessableEntityException(message ...string) HttpExceptionResponse {
	return CreateCustomException(422, message...)
}

func InternalServerErrorException(message ...string) HttpExceptionResponse {
	return CreateCustomException(500, message...)
}

func NotImplementedException(message ...string) HttpExceptionResponse {
	return CreateCustomException(501, message...)
}

func ImATeapotException(message ...string) HttpExceptionResponse {
	return CreateCustomException(418, message...)
}

func MethodNotAllowedException(message ...string) HttpExceptionResponse {
	return CreateCustomException(405, message...)
}

func BadGatewayException(message ...string) HttpExceptionResponse {
	return CreateCustomException(502, message...)
}

func ServiceUnavailableException(message ...string) HttpExceptionResponse {
	return CreateCustomException(503, message...)
}

func GatewayTimeoutException(message ...string) HttpExceptionResponse {
	return CreateCustomException(504, message...)
}

func PreconditionFailedException(message ...string) HttpExceptionResponse {
	return CreateCustomException(412, message...)
}
