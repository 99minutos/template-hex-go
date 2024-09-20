package errcodes

type MessageCode string

const (
	Generic500Error          MessageCode = "GENERIC_500_ERROR"
	Generic422Error          MessageCode = "GENERIC_422_ERROR"
	GenericInvalidParameters MessageCode = "GENERIC_INVALID_PARAMETERS"
	UnauthorizedRole         MessageCode = "UNAUTHORIZED_ROLE"
)

type ErrorCode struct {
	Code    int         `json:"code"`
	Message MessageCode `json:"message"`
}

func NewErrorCode(code int, message MessageCode) ErrorCode {
	return ErrorCode{
		Code:    code,
		Message: message,
	}
}

func (e ErrorCode) Error() string {
	return string(e.Message)
}
