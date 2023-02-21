package service

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strconv"
	"strings"
	"time"
)

func SendResponse(c *gin.Context, responseCode int, responseData interface{}) {

	c.Header("Content-Type", "application/json; charset=UTF-8")

	var doc []byte
	var err error
	bodyLen := 0

	if responseData != nil {
		doc, err = jsoniter.Marshal(responseData)
		if err != nil {
			log.Println("JsonMarshal Error [" + err.Error() + "]")
		}
		bodyLen = len(doc)
		c.Header("Content-Length", strconv.Itoa(bodyLen))
	} else {
		c.Header("Content-Length", "0")
	}

	corrId := c.Request.Header.Get("X-KM-Correlation-Id")
	caller := c.Request.Header.Get("X-KM-CALLER")
	userId := c.Request.Header.Get("X-Km-Userid")

	if caller == "" {
		caller = "CLIENT"
	}

	log.Printf("%s T[%s] U[%s] - [%s-RES]	 %s (Len:%d)", pkg.LogLevel(), corrId, userId, caller, string(doc), bodyLen)

	// GZIP 제외

	// Header 세팅되고 있는지 확인 필요
	c.Header("X-KM-TEMP", strconv.Itoa(responseCode))
	c.String(responseCode, string(doc))
}

func InitBaseResponseAttributes(handlerName string, response dto.BaseResponseMapper) {
	baseResponse := dto.BaseResponse{}

	handlerName = handlerName[strings.LastIndex(handlerName, ".")+1:]

	baseResponse.ExecutionStatus.StatusCodeData.SetSubject(handlerName)
	baseResponse.ExecutionStatus.SetStatusCodeData(dto.ErrNormal)
	baseResponse.ProcessingStart = time.Now().UnixMilli()

	response.BaseResponseMap(baseResponse)
}

func SetBaseResponseAttributes(errTasks map[string]dto.ErrorTask, response dto.BaseResponseMapper) {
	baseResponse := response.BaseResponseCopy()
	var reason, msg string

	if len(errTasks) == 0 {
		baseResponse.ExecutionStatus.SetStatusType("EXECUTED_SUCCESS")
	} else {
		baseResponse.ExecutionStatus.SetStatusType("FAILED")
		for _, v := range errTasks {
			msg = msg + "[" + v.ComponentCode() + "] " + v.RespError().GetMessage()
			reason = v.RespError().GetCode()
		}
		baseResponse.ExecutionStatus.StatusCodeData.SetReason(reason) // PIS.getVal() + ...
		baseResponse.ExecutionStatus.StatusCodeData.SetMessage(msg)
	}
	baseResponse.ProcessingEnd = time.Now().UnixMilli()

	response.BaseResponseMap(baseResponse)
}
