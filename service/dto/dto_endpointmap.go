package dto

import "sync"

type RequestEndpointMap map[string]RequestEndpointInfo

type EndpointMap struct {
	mu          *sync.RWMutex
	endpointMap RequestEndpointMap
}

func NewEndpointMap(endpointMap RequestEndpointMap) *EndpointMap {
	return &EndpointMap{mu: &sync.RWMutex{}, endpointMap: endpointMap}
}

func (r *EndpointMap) SetEndpointMap(m map[string]RequestEndpointInfo) {
	r.endpointMap = m
}

func (r EndpointMap) EndpointMap() (m map[string]RequestEndpointInfo) {
	m = r.endpointMap
	return
}

func (r *EndpointMap) Lock() {
	r.mu.Lock()
}

func (r *EndpointMap) UnLock() {
	r.mu.Unlock()
}

func (r *EndpointMap) Load(key string) (value RequestEndpointInfo, ok bool) {
	//TODO implement me
	r.mu.Lock()
	value, ok = r.endpointMap[key]
	r.mu.Unlock()
	return
}

func (r *EndpointMap) Store(key string, value RequestEndpointInfo) {
	//TODO implement me
	r.mu.Lock()
	if r.endpointMap == nil {
		r.endpointMap = make(map[string]RequestEndpointInfo)
	}
	r.endpointMap[key] = value
	r.mu.Unlock()
}
