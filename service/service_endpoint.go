package service

func InitPointParsInquiryEndpointMap(inputParam dto.PointParsInquiryRequestParam) (input dto.PointParsInquiryRequest, inputImpl dto.EndpointTasker, requestEndpointMap dto.RequestEndpointMap) {
	var reqEndpoint = configs.PointParsInquiryRequestEndpointMap()
	var endpoints = *dto.NewEndpointMap(reqEndpoint)

	input = *dto.NewPointParsInquiryRequest(inputParam, endpoints)

	inputImpl = &input
	inputImpl.SetEndpointMap(endpoints)
	requestEndpointMap = inputImpl.CopyRequestEndpointMap()

	return
}
