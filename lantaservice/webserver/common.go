package webserver

import (
	"lantaservice/usecase"
)

//func ServiceBuilder() {
//	usecase.NewPeriodService(entities.PeriodRepository)
//}

type HttpServer struct {
	CatalogService usecase.SiteRepository
}
