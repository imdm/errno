package errno

const (
	DefaultReasonDB             = "DB_ERROR"
	DefaultReasonRedis          = "REDIS_ERROR"
	DefaultReasonRPC            = "RPC_ERROR"
	DefaultReasonUnknown        = "INTERNAL_ERROR"
	DefaultReasonThirdParty     = "THIRD_PARTY_ERROR"
	DefaultReasonParamWrong     = "INVALID_PARAM"
	DefaultReasonSessionExpired = "SESSION_EXPIRED"
	DefaultReasonExisted        = "ALREADY_EXISTED"

	DefaultMessageDB             = "db error"
	DefaultMessageRedis          = "redis error"
	DefaultMessageRPC            = "internal rpc service error"
	DefaultMessageUnknown        = "internal error"
	DefaultMessageThirdParty     = "call third_party api failed"
	DefaultMessageParamWrong     = "invalid param"
	DefaultMessageSessionExpired = "session expired"
	DefaultMessageExisted        = "resource is already existed"
)

func Unknown(e ...error) *Error {
	msg := DefaultMessageUnknown
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonUnknown, msg)
}

func IsUnknown(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonUnknown
}

func DB(e ...error) *Error {
	msg := DefaultMessageDB
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonDB, msg)
}

func IsDB(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonDB
}

func Redis(e ...error) *Error {
	msg := DefaultMessageRedis
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonRedis, msg)
}

func IsRedis(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonRedis
}

func RPC(e ...error) *Error {
	msg := DefaultMessageRPC
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonRPC, msg)
}

func IsRPC(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonRPC
}

func ThirdParty(e ...error) *Error {
	msg := DefaultMessageThirdParty
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonThirdParty, msg)
}

func IsThirdParty(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonThirdParty
}
func ParamWrong(e ...error) *Error {
	msg := DefaultMessageParamWrong
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonParamWrong, msg)
}

func IsParamWrong(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonParamWrong
}
func SessionExpired(e ...error) *Error {
	msg := DefaultMessageSessionExpired
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonSessionExpired, msg)
}

func IsSessionExpired(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonSessionExpired
}

func Existed(e ...error) *Error {
	msg := DefaultMessageExisted
	if len(e) > 0 {
		msg = e[0].Error()
	}
	return InternalServer(DefaultReasonExisted, msg)
}

func IsExisted(e error) bool {
	err := FromError(e)
	return err.Code == 500 && err.Reason == DefaultReasonExisted
}
