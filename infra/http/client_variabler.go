package http

import "github.com/serverkona1/prepaid-card-inquiry-service/service/dto"

type ClientVariabler interface {
	GetQueryString() string
	GetPathVariable() string
	dto.IsSkippabler
}
