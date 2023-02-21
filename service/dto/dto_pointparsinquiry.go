package dto

import (
	"log"
)

type PointParsInquiryRequestParam struct {
	Par       string `json:"par" binding:"required"`
	ServiceId string `json:"serviceId"`
}

func (p PointParsInquiryRequestParam) GetPar() string {
	return p.Par
}

func (p PointParsInquiryRequestParam) Validate() bool {
	return !(p.Par == "")
}

type PointParsInquiryRequest struct {
	Param       PointParsInquiryRequestParam
	EndpointMap EndpointMap
	ErrorTasks  []ErrorTask
}

func NewPointParsInquiryRequest2(param PointParsInquiryRequestParam) *PointParsInquiryRequest {
	return &PointParsInquiryRequest{Param: param}
}

func NewPointParsInquiryRequest(param PointParsInquiryRequestParam, endpointMap EndpointMap) *PointParsInquiryRequest {
	return &PointParsInquiryRequest{Param: param, EndpointMap: endpointMap}
}

func (p PointParsInquiryRequest) RequestEndpointMapper(endPoint *RequestEndpointInfo) {
	//TODO implement me
	log.Println("PointParsInquiryRequest RequestEndpointMapper start")
	switch endPoint.TaskId() {
	case "IAS_BLC":
		IasBalanceListRequestMap(p.ParList(), endPoint)
		p.SetRequestEndpointMapEntity(*endPoint)
	case "IAS_CARD":
		IasBalanceListRequestMap(p.ParList(), endPoint)
		p.SetRequestEndpointMapEntity(*endPoint)
	case "IAS_CARD2":
		IasBalanceListRequestMap(p.ParList(), endPoint)
		p.SetRequestEndpointMapEntity(*endPoint)
	}
}

func (p *PointParsInquiryRequest) SetEndpointMap(endpointMap EndpointMap) {
	//TODO implement me
	p.EndpointMap = endpointMap
}

func (p PointParsInquiryRequest) GetRequestEndpointMap() RequestEndpointMap {
	//TODO implement me
	return p.EndpointMap.endpointMap
}

func (p PointParsInquiryRequest) CopyRequestEndpointMap() RequestEndpointMap {
	//TODO implement me
	endpointMap := RequestEndpointMap{}
	p.EndpointMap.mu.Lock()
	for k, v := range p.EndpointMap.endpointMap {
		endpointMap[k] = v
	}
	p.EndpointMap.mu.Unlock()
	return endpointMap
}

func (p *PointParsInquiryRequest) SetRequestEndpointMapEntity(endPoint RequestEndpointInfo) {
	//TODO implement me
	p.EndpointMap.Store(endPoint.TaskId(), endPoint)
}

//func (p *PointParsInquiryRequest) SetErrorTasks(taskID string, errMsg ResponseError) {
//	//TODO implement me
//	//p.Tasks.ErrorTasks[taskID] = errMsg
//	task := ErrorTask{}
//	task.SetComponentCode(taskID)
//	task.SetRespError(errMsg)
//	p.ErrorTasks = append(p.ErrorTasks, task)
//}
//
//func (p PointParsInquiryRequest) GetErrorTasks() []ErrorTask {
//	//TODO implement me
//	return p.ErrorTasks
//}

//func (p *PointParsInquiryRequest) SetDoneTasks(taskID string) {
//	//TODO implement me
//	p.Tasks.DoneTasks[taskID] = taskID
//}
//
//func (p PointParsInquiryRequest) GetDoneTasks() map[string]string {
//	//TODO implement me
//	return p.Tasks.DoneTasks
//}

func (p PointParsInquiryRequest) ParList() []string {
	var parList []string
	parList = append(parList, p.Param.Par)
	return parList
}

type PointParsInquiryResponse struct {
	// Common
	FunctionCallIdentifier int64           `json:"functionCallIdentifier"`
	ExecutionStatus        ExecutionStatus `json:"executionStatus"`
	ValidityPeriod         int64           `json:"validityPeriod"`
	ProcessingStart        int64           `json:"processingStart"`
	ProcessingEnd          int64           `json:"processingEnd"`
	//
	CardPointInfo CardPointInfo `json:"cardPointInfo"`
}

func (p PointParsInquiryResponse) BaseResponseCopy() (response BaseResponse) {
	//TODO implement me
	response.FunctionCallIdentifier = p.FunctionCallIdentifier
	response.ExecutionStatus = p.ExecutionStatus
	response.ValidityPeriod = p.ValidityPeriod
	response.ProcessingStart = p.ProcessingStart
	response.ProcessingEnd = p.ProcessingEnd

	return response
}

func (p *PointParsInquiryResponse) BaseResponseMap(response BaseResponse) {
	//TODO implement me
	p.FunctionCallIdentifier = response.FunctionCallIdentifier
	p.ExecutionStatus = response.ExecutionStatus
	p.ValidityPeriod = response.ValidityPeriod
	p.ProcessingStart = response.ProcessingStart
	p.ProcessingEnd = response.ProcessingEnd
}

func NewPointParsInquiryResponse() *PointParsInquiryResponse {
	return &PointParsInquiryResponse{CardPointInfo: CardPointInfo{PointInfos: []PointInfo{}, CompositePaymentInfos: []CompositePaymentInfo{}}}
}

func (p PointParsInquiryResponse) IsFollowingSkippable() bool {
	//TODO implement me
	return false
}

type PointInfo struct {
	PolicyType       string `json:"policyType"`
	PolicyId         string `json:"policyId,omitempty"`
	PolicyName       string `json:"policyName,omitempty"`
	Balance          int64  `json:"balance"`
	HaveGot          bool   `json:"haveGot"`
	IsCustomUsed     bool   `json:"isCustomUsed,omitempty"`
	IsAutoUsed       bool   `json:"isAutoUsed,omitempty"`
	EnableCustomUsed bool   `json:"enableCustomUsed,omitempty"`
}

type CompositePaymentInfo struct {
	PolicyId   string `json:"policyId"`
	PolicyName string `json:"policyName,omitempty"`
	Priority   int    `json:"-"`
}

type CardPointInfo struct {
	Par                   string                 `json:"-"`
	Balance               int64                  `json:"balance"`
	PointInfos            []PointInfo            `json:"pointInfos"`
	CompositePaymentInfos []CompositePaymentInfo `json:"compositePaymentInfos"`
}

func (c CardPointInfo) GetPar() string {
	return c.Par
}

func (c *CardPointInfo) SetBalance(balance int64) {
	c.Balance = balance
}
