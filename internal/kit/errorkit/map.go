package errorkit

type errMap struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`

	Err error `json:"-"`
}

type errMapping struct {
	UnexpectedError errMap
	OTPNotFound     errMap
}

var errorMapping = errMapping{
	UnexpectedError: errMap{"unknown_err", "Unexpected Error", nil},
	OTPNotFound:     errMap{"otp_not_found", "OTP Not Found", nil},
}

func ErrMap() errMapping {
	return errorMapping
}
