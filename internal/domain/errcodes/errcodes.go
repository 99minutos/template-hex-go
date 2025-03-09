package errcodes

type (
	MessageCode string
)

const (
	Generic500Error          MessageCode = "GENERIC_500_ERROR"
	Generic422Error          MessageCode = "GENERIC_422_ERROR"
	GenericInvalidParameters MessageCode = "GENERIC_INVALID_PARAMETERS"
	UnauthorizedRole         MessageCode = "UNAUTHORIZED_ROLE"

	// Example Codes
	ExampleNotFound          MessageCode = "EXAMPLE_NOT_FOUND"
	ExampleCodeAlreadyExists MessageCode = "EXAMPLE_CODE_ALREADY_EXISTS"
)

type ErrorCode struct {
	Code        int         `json:"code"`
	Message     MessageCode `json:"message"`
	Description string      `json:"description"`
}

func NewError(code int, message MessageCode) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: message,
	}
}

func NewErrorWithDescription(code int, message MessageCode, description string) *ErrorCode {
	return &ErrorCode{
		Code:        code,
		Message:     message,
		Description: description,
	}
}

func (e *ErrorCode) Error() string {
	return string(e.Message)
}
