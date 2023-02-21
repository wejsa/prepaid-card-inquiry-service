package dto

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ResponseError) GetCode() string {
	return e.Code
}

func (e ResponseError) GetMessage() string {
	return e.Message
}

func (e *ResponseError) SetMessage(message string) {
	e.Message = message
}

// Errors
var (
	ErrNormal           = ResponseError{Code: "000", Message: "Normal"}
	ErrConnectionFailed = ResponseError{Code: "057", Message: "External Components API Call Failure"}
	ErrParamInvalid     = ResponseError{Code: "901", Message: "Invalid Request Parameter"}
)
