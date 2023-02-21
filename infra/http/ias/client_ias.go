package ias

type BaseResponse struct {
	ExecutionStatus ExecutionStatus `json:"executionStatus"`
}

type ExecutionStatus struct {
	StatusType     string         `json:"statusType"`
	StatusCodeData StatusCodeData `json:"statusCodeData"`
}

type StatusCodeData struct {
	Subject string `json:"subject"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}
