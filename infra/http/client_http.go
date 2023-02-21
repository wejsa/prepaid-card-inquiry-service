package http

import (
	"bytes"
	"github.com/serverkona1/prepaid-card-inquiry-service/service/dto"
	"io"
	"log"
	"net/http"
	"time"
)

func ServerHTTP(info *dto.RequestEndpointInfo) error {

	var respBody []byte
	var additionalString = ""
	requestDto := info.RequestDto().(ClientVariabler)

	switch info.Method() {
	case "GET":
		additionalString += requestDto.GetPathVariable() + requestDto.GetQueryString()
		fullPath := info.Url() + additionalString

		log.Println("GET", fullPath)
		req, err := http.NewRequest("GET", fullPath, nil)
		if err != nil {
			log.Println("NewRequest Error : ", err.Error())
			return err // 오류 처리 필요
		}
		// Add Header

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("client.Get Error : " + err.Error())
			return err // 오류 처리 필요
		}
		defer resp.Body.Close()

		respBody, _ = io.ReadAll(resp.Body)
		log.Println("[" + info.TaskId() + "] resp [" + string(respBody) + "]")

	case "POST":
		additionalString += requestDto.GetPathVariable() + requestDto.GetQueryString()
		fullPath := info.Url() + additionalString
		reqBody := bytes.NewBufferString(info.Request())

		log.Println("POST", fullPath, info.Request())
		req, err := http.NewRequest("POST", fullPath, reqBody)
		if err != nil {
			log.Println("NewRequest Error : ", err.Error())
			return err // 오류 처리 필요
		}
		// Add Header
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{Timeout: 10 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("client.Post Error : " + err.Error())
			return err // 오류 처리 필요
		}
		defer resp.Body.Close()

		respBody, _ = io.ReadAll(resp.Body)
		log.Println("[" + info.TaskId() + "] resp [" + string(respBody) + "]")
	}

	info.SetResponse(string(respBody))

	return nil
}
