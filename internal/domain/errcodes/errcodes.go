package errcodes

type MessageCode string
type DescriptionFromCode string

const (
	Generic500Error          MessageCode = "GENERIC_500_ERROR"
	Generic422Error          MessageCode = "GENERIC_422_ERROR"
	GenericInvalidParameters MessageCode = "GENERIC_INVALID_PARAMETERS"
	UnauthorizedRole         MessageCode = "UNAUTHORIZED_ROLE"

	// Carrier Codes
	CarrierNotFound          MessageCode = "CARRIER_NOT_FOUND"
	CarrierCodeAlreadyExists MessageCode = "CARRIER_CODE_ALREADY_EXISTS"
)

const (
	CarrierCodeAlreadyExistsDesc DescriptionFromCode = "El código de la transportadora ya existe"
)

type ErrorCode struct {
	Code        int                 `json:"code"`
	Message     MessageCode         `json:"message"`
	Description DescriptionFromCode `json:"description"`
}

func NewError(code int, message MessageCode) *ErrorCode {
	return &ErrorCode{
		Code:    code,
		Message: message,
	}
}

func NewErrorWithDescription(code int, message MessageCode, description DescriptionFromCode) *ErrorCode {
	return &ErrorCode{
		Code:        code,
		Message:     message,
		Description: description,
	}
}

func (e *ErrorCode) Error() string {
	return string(e.Message)
}
