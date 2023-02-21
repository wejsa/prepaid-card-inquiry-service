package dto

type BaseResponseMapper interface {
	BaseResponseMap(response BaseResponse)
	BaseResponseCopy() (response BaseResponse)
}

type BaseResponse struct {
	FunctionCallIdentifier int64
	ExecutionStatus        ExecutionStatus
	ValidityPeriod         int64
	ProcessingStart        int64
	ProcessingEnd          int64
}

type ExecutionStatus struct {
	StatusCodeData StatusCodeData `json:"statusCodeData"`
	StatusType     string         `json:"statusType"`
}

func (e ExecutionStatus) GetStatusCodeData() StatusCodeData {
	return e.StatusCodeData
}

func (e ExecutionStatus) GetStatusType() string {
	return e.StatusType
}

func (e *ExecutionStatus) SetStatusCodeData(errMsg ResponseError) {
	e.StatusCodeData.Reason = errMsg.Code
	e.StatusCodeData.Message = errMsg.Message
}

func (e *ExecutionStatus) SetStatusType(statusType string) {
	e.StatusType = statusType
}

func NewExecutionStatus(statusCodeData StatusCodeData, statusType string) *ExecutionStatus {
	return &ExecutionStatus{StatusCodeData: statusCodeData, StatusType: statusType}
}

type StatusCodeData struct {
	Subject           string `json:"subject,omitempty"`
	SubjectIdentifier string `json:"subjectIdentifier,omitempty"`
	Reason            string `json:"reason,omitempty"`
	Message           string `json:"message,omitempty"`
}

func (s *StatusCodeData) SetSubject(Subject string) {
	s.Subject = Subject
}

func (s *StatusCodeData) SetSubjectIdentifier(SubjectIdentifier string) {
	s.SubjectIdentifier = SubjectIdentifier
}

func (s *StatusCodeData) SetReason(Reason string) {
	s.Reason = Reason
}

func (s *StatusCodeData) SetMessage(Message string) {
	s.Message = Message
}

func (s StatusCodeData) GetReason() string {
	return s.Reason
}
