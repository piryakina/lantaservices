package webserver

import "net/http"

//func ServiceBuilder() {
//	usecase.NewPeriodService(entities.PeriodRepository)
//}
//
//type HttpServer struct {
//	CatalogService usecase.SiteRepository
//}

func Front(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../front/dist/lantafront/index.html")
}
