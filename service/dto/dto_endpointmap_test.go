package dto

import (
	"github.com/serverkona1/prepaid-card-inquiry-service/infra/http/ias"
	"log"
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {

	endpointMap := EndpointMap{}

	endpointMap.mu = &sync.RWMutex{}
	endpointMap.SetEndpointMap(initTestMap())

	m := endpointMap.EndpointMap()

	tmpMap := make(map[string]RequestEndpointInfo)
	endpointMap.mu.Lock()
	for k, v := range endpointMap.EndpointMap() {
		tmpMap[k] = v
	}
	endpointMap.mu.Unlock()

	done := make(chan bool)
	for i := 0; i < 10000; i++ {
		for k, v := range tmpMap {
			go func(key string, value RequestEndpointInfo, c chan<- bool) {
				//if value, ok := endpointMap.Load(key); ok {
				info := value
				info.request = value.taskID + "-" + value.method
				endpointMap.Store(key, info)
				//}

				c <- true
			}(k, v, done)
		}

		for i := 0; i < len(m); i++ {
			<-done
		}
	}

	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>> done..")

	for k, v := range m {
		log.Println("key:", k, ", value:", v)
	}
}

func initTestMap() (m map[string]RequestEndpointInfo) {
	m = make(map[string]RequestEndpointInfo)
	m["IAS"] = *NewRequestEndpointInfo("IAS", "IAS", "http://127.0.0.1:5000/sample1", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS0"] = *NewRequestEndpointInfo("KPS0", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS1"] = *NewRequestEndpointInfo("KPS1", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS2"] = *NewRequestEndpointInfo("KPS2", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS3"] = *NewRequestEndpointInfo("KPS3", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS4"] = *NewRequestEndpointInfo("KPS4", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS5"] = *NewRequestEndpointInfo("KPS5", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS6"] = *NewRequestEndpointInfo("KPS6", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS7"] = *NewRequestEndpointInfo("KPS7", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS8"] = *NewRequestEndpointInfo("KPS8", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["KPS9"] = *NewRequestEndpointInfo("KPS9", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["IAS1"] = *NewRequestEndpointInfo("IAS1", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)
	m["IAS2"] = *NewRequestEndpointInfo("IAS2", "KPS", "http://127.0.0.1:5000/sample2", false, "POST", ias.BalanceListRequest{}, ias.BalanceListResponse{}, "Y", 1)

	return
}
