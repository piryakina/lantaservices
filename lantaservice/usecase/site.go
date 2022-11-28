package usecase

type ServiceSite struct {
	catalog SiteRepository
}

type SiteRepository interface {
	PeriodServer
	LoginServer
	NewsServer
	UserServer
	//UserRepository
	//TreeRepository
}

//
//// NewServiceSite constructor for catalogstorage service
//func NewServiceSite(storage SiteRepository) *ServiceSite {
//	return &ServiceSite{catalog: storage}
//}
//
//// SiteServer use case for catalogstorage
//type SiteServer interface {
//}
