package configs

import (
	"github.com/serverkona1/prepaid-card-inquiry-service/infra/http/ias"
	"github.com/serverkona1/prepaid-card-inquiry-service/service/dto"
)

func PointParsInquiryRequestEndpointMap() (endpointMap dto.RequestEndpointMap) {

	endpointMap = make(dto.RequestEndpointMap)
	endpointMap["IAS_BLC"] = *dto.NewRequestEndpointInfo("IAS_BLC", "IAS", "http://127.0.0.1:5000/sample1", false, "POST",
		ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	endpointMap["IAS_CARD"] = *dto.NewRequestEndpointInfo("IAS_CARD", "IAS", "http://127.0.0.1:5000/sample2", false, "POST",
		ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	endpointMap["IAS_CARD"] = *dto.NewRequestEndpointInfo("IAS_CARD", "IAS", "http://127.0.0.1:5000/sample2", false, "POST",
		ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	//endpointMap["IAS"] = *dto.NewRequestEndpointInfo("IAS", "IAS", "http://118.33.122.28:10100/issuer-authorization-system-1.0/api/getBalanceList",
	//	false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	//endpointMap["KPS"] = *model.NewRequestEndpointInfo("KPS", "KPS", "http://127.0.0.1:5000/sample2",
	//	false, "GET", dto.KpsHoldingCardPointListRequest{}, dto.KpsHoldingCardPointListResponse{}, 1)

	return
}
