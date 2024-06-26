package errorcode

import "google.golang.org/grpc/codes"

const (
	InternalServerError = "INTERNAL_SERVER_ERROR"
	InternalServerPanic = "INTERNAL_SERVER_PANIC"
	AuthenticationError = "AUTHENTICATION_ERROR"
	AuthorizationError  = "AUTHORIZATION_ERROR"
	BadRequestError     = "BAD_REQUEST_ERROR"
	ValidationError     = "VALIDATION_ERROR"
	NotFoundError       = "NOT_FOUND_ERROR"
)

// ErrorsMap maps error codes with grpc's error codes, which define http status codes
// https://cloud.yandex.com/en/docs/api-design-guide/concepts/errors
var ErrorsMap = map[string]codes.Code{
	InternalServerError: codes.Internal,
	InternalServerPanic: codes.Internal,
	AuthenticationError: codes.Unauthenticated,
	AuthorizationError:  codes.PermissionDenied,
	BadRequestError:     codes.InvalidArgument,
	ValidationError:     codes.InvalidArgument,
	NotFoundError:       codes.NotFound,
}
