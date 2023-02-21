package dto

type EndpointTasker interface {
	RequestEndpointMapper(requestEndpointInfo *RequestEndpointInfo)
	SetEndpointMap(endpointMap EndpointMap)
	GetRequestEndpointMap() RequestEndpointMap
	CopyRequestEndpointMap() RequestEndpointMap
	SetRequestEndpointMapEntity(requestEndpointInfo RequestEndpointInfo)

	//SetErrorTasks(taskID string, responseError ResponseError)
	//GetErrorTasks() []ErrorTask
	//SetDoneTasks(taskID string)
	//GetDoneTasks() map[string]string
}
