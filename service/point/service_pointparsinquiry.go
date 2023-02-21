package point

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/serverkona1/prepaid-card-inquiry-service/infra/http/ias"
	"github.com/serverkona1/prepaid-card-inquiry-service/service"
	"github.com/serverkona1/prepaid-card-inquiry-service/service/dto"
)

func PostPointParsInquiryResponseMap(response dto.PointParsInquiryResponse, respEndpoints dto.RequestEndpointMap, par string) (result dto.PointParsInquiryResponse) {
	result = response
	result.CardPointInfo.Par = par
	endpoints := respEndpoints

	doneTasks, errorTasks := .GetResultByTasks(endpoints)

	for k, v := range endpoints {
		if _, ok := doneTasks[k]; ok {
			switch k {
			case "IAS_BLC":
				iasBalanceListResponseToPointParsInquiryResponseMap(v.Response(), &result)
			case "IAS_CARD":
				iasBalanceListResponseToPointParsInquiryResponseMap(v.Response(), &result)
			case "IAS_CARD2":
				iasBalanceListResponseToPointParsInquiryResponseMap(v.Response(), &result)
			}
		}
	}

	//errorTasks := service.GetErrorTasks(endpoints)
	service.SetBaseResponseAttributes(errorTasks, &result)

	//log.Println("resp:", result)
	return result
}

func iasBalanceListResponseToPointParsInquiryResponseMap(respBody string, result *dto.PointParsInquiryResponse) {
	responseJSON := ias.BalanceListResponse{}
	jsoniter.Unmarshal([]byte(respBody), &responseJSON)

	for _, v := range responseJSON.BalanceList {
		if result.CardPointInfo.GetPar() == v.GetPar() {
			result.CardPointInfo.SetBalance(v.GetBalance())
		}
	}
}
