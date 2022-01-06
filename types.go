package errno

import "net/http"

const (
	DefaultReasonBadRequest         = "INVALID_PARAM"
	DefaultReasonUnauthorized       = "UNAUTHORIZED"
	DefaultReasonForbidden          = "FORBIDDEN"
	DefaultReasonNotFound           = "NOT_FOUND"
	DefaultReasonConflict           = "CONFLICT"
	DefaultReasonTooManyRequests    = "TOO_MANY_REQUESTS"
	DefaultReasonInternalServer     = "INTERNAL_SERVER"
	DefaultReasonNotImplemented     = "NOT_IMPLEMENTED"
	DefaultReasonServiceUnavailable = "SERVICE_UNAVAILABLE"
	DefaultReasonGatewayTimeout     = "GATEWAY_TIMEOUT"
	DefaultReasonClientClosed       = "CLIENT_CLOSED"

	DefaultMessageBadRequest         = "invalid param"
	DefaultMessageUnauthorized       = "unauthorized"
	DefaultMessageForbidden          = "forbidden"
	DefaultMessageNotFound           = "not found"
	DefaultMessageConflict           = "conflict"
	DefaultMessageTooManyRequests    = "too many requests"
	DefaultMessageInternalServer     = "internal server"
	DefaultMessageNotImplemented     = "not implemented"
	DefaultMessageServiceUnavailable = "service unavailable"
	DefaultMessageGateWayTimeout     = "gateway timeout"
	DefaultMessageClientClosed       = "client closed"
)

// BadRequest new BadRequest error that is mapped to a 400 response.
func BadRequest(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonBadRequest
	}
	if message == "" {
		message = DefaultMessageBadRequest
	}
	return Newf(http.StatusBadRequest, reason, message)
}

// IsBadRequest determines if err is an error which indicates a BadRequest error.
// It supports wrapped errors.
func IsBadRequest(err error) bool {
	return Code(err) == http.StatusBadRequest
}

// Unauthorized new Unauthorized error that is mapped to a 401 response.
func Unauthorized(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonUnauthorized
	}
	if message == "" {
		message = DefaultMessageUnauthorized
	}
	return Newf(http.StatusUnauthorized, reason, message)
}

// IsUnauthorized determines if err is an error which indicates an Unauthorized error.
// It supports wrapped errors.
func IsUnauthorized(err error) bool {
	return Code(err) == http.StatusUnauthorized
}

// Forbidden new Forbidden error that is mapped to a 403 response.
func Forbidden(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonForbidden
	}
	if message == "" {
		message = DefaultMessageForbidden
	}
	return Newf(http.StatusForbidden, reason, message)
}

// IsForbidden determines if err is an error which indicates a Forbidden error.
// It supports wrapped errors.
func IsForbidden(err error) bool {
	return Code(err) == http.StatusForbidden
}

// NotFound new NotFound error that is mapped to a 404 response.
func NotFound(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonNotFound
	}
	if message == "" {
		message = DefaultMessageNotFound
	}
	return Newf(http.StatusNotFound, reason, message)
}

// IsNotFound determines if err is an error which indicates an NotFound error.
// It supports wrapped errors.
func IsNotFound(err error) bool {
	return Code(err) == http.StatusNotFound
}

// Conflict new Conflict error that is mapped to a 409 response.
func Conflict(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonConflict
	}
	if message == "" {
		message = DefaultMessageConflict
	}
	return Newf(http.StatusConflict, reason, message)
}

// IsConflict determines if err is an error which indicates a Conflict error.
// It supports wrapped errors.
func IsConflict(err error) bool {
	return Code(err) == http.StatusConflict
}

// TooManyRequests new TooManyRequests error that is mapped to an HTTP 429 response.
func TooManyRequests(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonTooManyRequests
	}
	if message == "" {
		message = DefaultMessageTooManyRequests
	}
	return Newf(http.StatusTooManyRequests, reason, message)
}

// IsTooManyRequests determines if err is an error which indicates a TooManyRequests error.
// It supports wrapped errors.
func IsTooManyRequests(err error) bool {
	return Code(err) == http.StatusTooManyRequests
}

// InternalServer new InternalServer error that is mapped to a 500 response.
func InternalServer(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonInternalServer
	}
	if message == "" {
		message = DefaultMessageInternalServer
	}
	return Newf(http.StatusInternalServerError, reason, message)
}

// IsInternalServer determines if err is an error which indicates an Internal error.
// It supports wrapped errors.
func IsInternalServer(err error) bool {
	return Code(err) == http.StatusInternalServerError
}

// NotImplemented new NotImplemented error that is mapped to a 500 response.
func NotImplemented(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonNotImplemented
	}
	if message == "" {
		message = DefaultMessageNotImplemented
	}
	return Newf(http.StatusNotImplemented, reason, message)
}

// IsNotImplemented determines if err is an error which indicates an NotImplemented error.
// It supports wrapped errors.
func IsNotImplemented(err error) bool {
	return Code(err) == http.StatusNotImplemented
}

// ServiceUnavailable new ServiceUnavailable error that is mapped to an HTTP 503 response.
func ServiceUnavailable(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonServiceUnavailable
	}
	if message == "" {
		message = DefaultMessageServiceUnavailable
	}
	return Newf(http.StatusServiceUnavailable, reason, message)
}

// IsServiceUnavailable determines if err is an error which indicates an Unavailable error.
// It supports wrapped errors.
func IsServiceUnavailable(err error) bool {
	return Code(err) == http.StatusServiceUnavailable
}

// GatewayTimeout new GatewayTimeout error that is mapped to an HTTP 504 response.
func GatewayTimeout(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonGatewayTimeout
	}
	if message == "" {
		message = DefaultMessageGateWayTimeout
	}
	return Newf(http.StatusGatewayTimeout, reason, message)
}

// IsGatewayTimeout determines if err is an error which indicates a GatewayTimeout error.
// It supports wrapped errors.
func IsGatewayTimeout(err error) bool {
	return Code(err) == http.StatusGatewayTimeout
}

// ClientClosed new ClientClosed error that is mapped to an HTTP 499 response.
func ClientClosed(reason, message string) *Error {
	if reason == "" {
		reason = DefaultReasonClientClosed
	}
	if message == "" {
		message = DefaultMessageClientClosed
	}
	return Newf(499, reason, message)
}

// IsClientClosed determines if err is an error which indicates a ClientClosed error.
// It supports wrapped errors.
func IsClientClosed(err error) bool {
	return Code(err) == 499
}
