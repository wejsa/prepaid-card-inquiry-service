package service

import (
	jsoniter "github.com/json-iterator/go"
	"log"
)

func WorkerNoGoroutine(input dto.EndpointTasker, info dto.RequestEndpointInfo) {

	input.RequestEndpointMapper(&info)

	process(&info)

	done(input, &info)
}

// Worker is common task
func Worker(tasker dto.EndpointTasker, info dto.RequestEndpointInfo, c chan<- bool) {

	tasker.RequestEndpointMapper(&info)

	process(&info)

	done(tasker, &info)

	c <- true
}

// process is service communication
func process(info *dto.RequestEndpointInfo) {
	if err := isSkippable(info); err == true {
		return
	}

	// 컴포넌트 통신
	err := pishttp.ServerHTTP(info)
	if err != nil {
		info.SetErrorMessage(dto.ErrConnectionFailed)
	}
	return
}

// done processing the completion of the task
func done(tasker dto.EndpointTasker, info *dto.RequestEndpointInfo) {
	if isFailed(info) {
		log.Println("[" + info.TaskId() + "] isFailed end")
	} else {
		log.Println("[" + info.TaskId() + "] Success task end")
	}

	tasker.SetRequestEndpointMapEntity(*info)
}

// isSkippable checks the skip Endpoint
func isSkippable(info *dto.RequestEndpointInfo) bool {
	requestDto := info.RequestDto().(dto.IsSkippabler)
	responseDto := info.ResponseDto().(dto.IsFollowingSkippabler)
	if requestDto.IsSkippable() ||
		(info.IsSkippable() && responseDto.IsFollowingSkippable()) {
		return true
	}
	return false
}

// isFailed checks result code of Endpoint response data
func isFailed(info *dto.RequestEndpointInfo) bool {
	// 컴포넌트별로 baseResponse 가 상이하여 각각 생성하여 응답 결과 체크 필요
	switch info.ComponentCode() {
	case "IAS":
		response := ias.BaseResponse{}
		jsoniter.Unmarshal([]byte(info.Response()), &response)
		log.Println("["+info.TaskId()+"] response:", response, "statusType:"+response.ExecutionStatus.StatusType)
		if response.ExecutionStatus.StatusType != "EXECUTED_SUCCESS" {
			responseError := dto.ErrConnectionFailed
			if response.ExecutionStatus.StatusCodeData.Message != "" {
				responseError.SetMessage(response.ExecutionStatus.StatusCodeData.Message)
			}
			info.SetErrorMessage(responseError)
			log.Println("["+info.TaskId()+"] ErrorMessage:", info.ErrorMessage())
			return true
		} else {
			return false
		}
	}

	return false
}
