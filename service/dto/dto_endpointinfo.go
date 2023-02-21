package dto

type RequestEndpointInfo struct {
	taskID        string
	componentCode string
	url           string
	isSkippable   bool
	method        string
	requestDto    interface{}
	responseDto   interface{}
	baseResponse  BaseResponse
	request       string
	response      string
	isSync        string
	priority      int
	errResponse   ResponseError
}

func (r *RequestEndpointInfo) SetTaskID(taskID string) {
	r.taskID = taskID
}

func (r *RequestEndpointInfo) SetComponentCode(componentCode string) {
	r.componentCode = componentCode
}

func (r *RequestEndpointInfo) SetUrl(url string) {
	r.url = url
}

func (r *RequestEndpointInfo) SetIsSkippable(isSkippable bool) {
	r.isSkippable = isSkippable
}

func (r *RequestEndpointInfo) SetMethod(method string) {
	r.method = method
}

func (r *RequestEndpointInfo) SetRequestDto(requestDto interface{}) {
	r.requestDto = requestDto
}

func (r *RequestEndpointInfo) SetResponseDto(responseDto interface{}) {
	r.responseDto = responseDto
}

func (r *RequestEndpointInfo) SetBaseResponse(baseResponse BaseResponse) {
	r.baseResponse = baseResponse
}

func (r *RequestEndpointInfo) SetRequest(request string) {
	r.request = request
}

func (r *RequestEndpointInfo) SetResponse(response string) {
	r.response = response
}

func (r *RequestEndpointInfo) SetIsSync(isSync string) {
	r.isSync = isSync
}

func (r *RequestEndpointInfo) SetPriority(priority int) {
	r.priority = priority
}

func (r *RequestEndpointInfo) SetErrorMessage(errResponse ResponseError) {
	r.errResponse = errResponse
}

func (r RequestEndpointInfo) TaskId() string {
	return r.taskID
}

func (r RequestEndpointInfo) ComponentCode() string {
	return r.componentCode
}

func (r RequestEndpointInfo) Url() string {
	return r.url
}

func (r RequestEndpointInfo) IsSkippable() bool {
	return r.isSkippable
}

func (r RequestEndpointInfo) Method() string {
	return r.method
}

func (r RequestEndpointInfo) RequestDto() interface{} {
	return r.requestDto
}

func (r RequestEndpointInfo) ResponseDto() interface{} {
	return r.responseDto
}

func (r RequestEndpointInfo) BaseResponse() BaseResponse {
	return r.baseResponse
}

func (r RequestEndpointInfo) Request() string {
	return r.request
}

func (r RequestEndpointInfo) Response() string {
	return r.response
}

func (r RequestEndpointInfo) IsSync() string {
	return r.isSync
}

func (r RequestEndpointInfo) Priority() int {
	return r.priority
}

func (r RequestEndpointInfo) ErrorMessage() ResponseError {
	return r.errResponse
}

func NewRequestEndpointInfo(key string, componentCode string, url string, isSkippable bool, method string,
	requestDto interface{}, responseDto interface{}, isSync string, priority int) *RequestEndpointInfo {
	return &RequestEndpointInfo{
		taskID:        key,
		componentCode: componentCode,
		url:           url,
		isSkippable:   isSkippable,
		method:        method,
		requestDto:    requestDto,
		responseDto:   responseDto,
		isSync:        isSync,
		priority:      priority,
	}
}
