package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/serverkona1/prepaid-card-inquiry-service/api/routes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutePointParsInquiry(t *testing.T) {

	router := routes.NewRouter()

	w := httptest.NewRecorder()

	test1(w, router, t)

}

func test1(w *httptest.ResponseRecorder, r *gin.Engine, t *testing.T) {

	//reqBody := bytes.NewBufferString("[{\"par\":\"par123455\", \"serviceId\":\"S00000\", \"serviceId2\":\"S000002\"}, {\"par\":\"par123455\", \"serviceId\":\"S00000\", \"serviceId2\":\"S000002\"}]")
	reqBody := bytes.NewBufferString("{\"par\":\"Q152993AF0FF204D58CED6AEE72\", \"serviceId\":\"S00000\"}")
	req, _ := http.NewRequest("POST", "/api/v3/point/pars/inquiry", reqBody)
	req.Header.Set("X-KM-Correlation-Id", "TID-01")
	req.Header.Set("X-Km-Userid", "U-0001")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("XM-ADD", "Go-test")
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
