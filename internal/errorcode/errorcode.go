package errorcode

import "google.golang.org/grpc/codes"

const (
	InternalServerError = "INTERNAL_SERVER_ERROR"
	InternalServerPanic = "INTERNAL_SERVER_PANIC"
	BadRequestError     = "BAD_REQUEST_ERROR"
	NotFoundError       = "NOT_FOUND_ERROR"
)

// ErrorsMap maps error codes with grpc's error codes, which define http status codes
// https://cloud.yandex.com/en/docs/api-design-guide/concepts/errors
var ErrorsMap = map[string]codes.Code{
	InternalServerError: codes.Internal,
	InternalServerPanic: codes.Internal,
	BadRequestError:     codes.InvalidArgument,
	NotFoundError:       codes.NotFound,
}
