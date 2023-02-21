package dto

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/serverkona1/prepaid-card-inquiry-service/infra/http/ias"
)

func IasBalanceListRequestMap(parList []string, endPoint *RequestEndpointInfo) {
	iasBalanceListRequest := ias.BalanceListRequest{}
	iasBalanceListRequest.SetParList(parList)

	doc, _ := jsoniter.Marshal(iasBalanceListRequest)

	endPoint.SetRequestDto(iasBalanceListRequest)
	endPoint.SetRequest(string(doc))
}
