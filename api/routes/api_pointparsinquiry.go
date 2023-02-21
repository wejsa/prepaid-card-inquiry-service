package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/serverkona1/prepaid-card-inquiry-service/service"
	"github.com/serverkona1/prepaid-card-inquiry-service/service/dto"
	"github.com/serverkona1/prepaid-card-inquiry-service/service/point"
	"log"
	"net/http"
)

func PointParsInquiry(c *gin.Context) {
	log.Println("PointParsInquiry start")
	var inputParam dto.PointParsInquiryRequestParam

	responseCode := http.StatusOK
	response := *dto.NewPointParsInquiryResponse()

	service.InitBaseResponseAttributes(c.HandlerName(), &response)

	err := c.BindJSON(&inputParam)
	if err != nil {
		log.Println(err.Error())
		response.ExecutionStatus.SetStatusCodeData(dto.ErrParamInvalid) // 변경
		service.SendResponse(c, http.StatusBadRequest, &response)
		return
	}

	is := inputParam.Validate()
	if is == false {
		log.Println(err.Error())
		response.ExecutionStatus.SetStatusCodeData(dto.ErrParamInvalid) // 변경
		service.SendResponse(c, http.StatusBadRequest, &response)
		return
	}

	var input dto.PointParsInquiryRequest
	var inputImpl dto.EndpointTasker
	var endpointMap dto.RequestEndpointMap

	input, inputImpl, endpointMap = service.InitPointParsInquiryEndpointMap(inputParam)

	// 우선순위 체크로직 조회 추가

	done := make(chan bool)
	for k, v := range endpointMap {
		log.Printf("%s start...", k)

		go service.Worker(inputImpl, v, done)
	}

	for i := 0; i < len(endpointMap); i++ {
		<-done
	}

	response = point.PostPointParsInquiryResponseMap(response, inputImpl.GetRequestEndpointMap(), input.Param.GetPar())

	service.SendResponse(c, responseCode, &response)
}
