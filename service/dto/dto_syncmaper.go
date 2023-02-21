package dto

type mapInterface interface {
	Load(key string) (RequestEndpointInfo, bool)
	Store(key string, value RequestEndpointInfo)
	Lock()
	UnLock()
}
