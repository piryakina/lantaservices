package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/code/assets"))))
	for _, route := range Routing() {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	router.NotFoundHandler = router.NewRoute().HandlerFunc(Index).GetHandler()

	return router
}

func Routing() []Route {
	return []Route{
		{"test", "GET", "/test", test},
		{"index", "GET", "/", Index},
		{"login", "POST", "/login", Login},
		{"logout", "GET", "/logout", Logout},
		//users
		{"add-user", "POST", "/admin/add-user", AddUser},
		{"add-user", "GET", "/user", GetUserRoleById},
		//{"edit-user", "POST", "/admin/edit-user", nil},
		{"add-sp", "POST", "/admin/add-sp", AddSp},
		//{"edit-sp", "POST", "/admin/edit-sp", nil},
		//отчетные периоды
		{"add-period", "POST", "/admin/add-period", AddNewPeriod},
		//{"edit-period", "POST", "/admin/edit-period", nil},
		{"get-period", "GET", "/period-now", GetPeriodNow},
		{"get-periods", "GET", "/all-period", GetAllPeriods},
		//news
		{"add-news", "POST", "/admin/add-news", AddNews},
		//{"edit-news", "POST", "/admin/edit-news", nil},
		{"get-news", "GET", "/news", GetNews},

		//services partner
		{"get-quality-and-process", "GET", "/sp-period", GetDataSpPeriodNow},
		{"add-quality-and-process", "POST", "/add/sp-period", AddDataSpPeriodNow},
		//add billing by sp
		{"upload-file", "POST", "/billings/upload", nil},
		//add score by sp
		{"upload-file", "POST", "/result/upload", nil},
		//add response by analytic or usp
		{"upload-file", "POST", "/response/upload", nil},
	}
}
