package dto

type ErrorTask struct {
	componentCode string
	respError     ResponseError
}

func NewErrorTask(componenctCode string, responseError ResponseError) *ErrorTask {
	return &ErrorTask{componentCode: componenctCode, respError: responseError}
}

func (e ErrorTask) ComponentCode() string {
	return e.componentCode
}

func (e ErrorTask) RespError() ResponseError {
	return e.respError
}

func (e *ErrorTask) SetComponentCode(code string) {
	e.componentCode = code
}

func (e *ErrorTask) SetRespError(respError ResponseError) {
	e.respError = respError
}
