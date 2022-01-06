package errno

import (
	"errors"
	"fmt"

	"github.com/imdm/errno/internal/httputil"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
	// UnknownReason is unknown reason for error info.
	UnknownReason = ""
	// SupportPackageIsVersion1 this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
)

//go:generate protoc -I. --go_out=paths=source_relative:. errors.proto

func (x *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s metadata = %v", x.Code, x.Reason, x.Message, x.Metadata)
}

// GRPCStatus returns the Status represented by se.
func (x *Error) GRPCStatus() *status.Status {
	s, _ := status.New(httputil.ToGRPCCode(int(x.Code)), x.Message).
		WithDetails(&errdetails.ErrorInfo{
			Reason:   x.Reason,
			Metadata: x.Metadata,
		})
	return s
}

// Is matches each error in the chain with the target value.
func (x *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Reason == x.Reason
	}
	return false
}

// CopyWithMetadata copies the error instance and returns a new one with an MD formed by the mapping of key, value.
func (x *Error) CopyWithMetadata(md map[string]string) *Error {
	err := proto.Clone(x).(*Error)
	err.Metadata = md
	return err
}

// CopyWithMessage copies the error instance and returns a new one with given message.
func (x *Error) CopyWithMessage(msg string) *Error {
	err := proto.Clone(x).(*Error)
	err.Message = msg
	return err
}

// CopyWithCauseError copies the error instance and returns a new one with given error.
func (x *Error) CopyWithCauseError(e error) *Error {
	err := proto.Clone(x).(*Error)
	err.Message = e.Error()
	return err
}

// New returns an error object for the code, message.
func New(code int, reason, message string) *Error {
	return &Error{
		Code:    int32(code),
		Message: message,
		Reason:  reason,
	}
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Errorf returns an error object for the code, message and error info.
func Errorf(code int, reason, format string, a ...interface{}) error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Code returns the http code for a particular error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 200
	}
	if se := FromError(err); err != nil {
		return int(se.Code)
	}
	return UnknownCode
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if se := FromError(err); err != nil {
		return se.Reason
	}
	return UnknownReason
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	gs, ok := status.FromError(err)
	if ok {
		ret := New(
			httputil.FromGRPCCode(gs.Code()),
			UnknownReason,
			gs.Message(),
		)
		for _, detail := range gs.Details() {
			switch d := detail.(type) {
			case *errdetails.ErrorInfo:
				ret.Reason = d.Reason
				return ret.CopyWithMetadata(d.Metadata)
			}
		}
		return ret
	}
	return New(UnknownCode, UnknownReason, err.Error())
}
