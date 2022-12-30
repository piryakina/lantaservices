package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
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
		//{"index", "GET", "/", Index},
		{"login", "POST", "/login", Login},
		{"logout", "GET", "/logout", Logout},
		//users
		{"add-user", "POST", "/admin/add-user", AddUser},
		{"check-user-unique-login", "POST", "/admin/checklogin", CheckLogin},
		{"add-user", "GET", "/role", GetUserRoleById},
		{"add-user", "GET", "/roles", GetRoles},
		{"add-user", "GET", "/sp", GetSPNameById},
		//{"edit-user", "POST", "/admin/edit-user", nil},
		{"add-sp", "POST", "/admin/add-sp", AddSp},
		//{"edit-sp", "POST", "/admin/edit-sp", nil},

		//usp
		{"upload-file", "POST", "/usp/sla", UploadSLA},
		{"download-file", "GET", "/usp/sla/download", DownloadSLA},
		//отчетные периоды
		{"add-period", "POST", "/admin/add-period", AddNewPeriod},
		//{"edit-period", "POST", "/admin/edit-period", nil},
		{"get-period", "GET", "/period-now", GetPeriodNow},
		{"get-periods", "GET", "/all-period", GetAllPeriods},
		//news
		{"add-news", "POST", "/admin/add-news", AddNews},
		//{"edit-news", "POST", "/admin/edit-news", nil},
		{"get-news", "GET", "/news", GetNews},
		{"get-attach", "GET", "/news-img", GetImg},

		//status file
		{"get-status-file", "GET", "/status", GetStatuses},
		{"set-status-file", "GET", "/set-status", SetStatusFile},

		//services partner
		{"get-quality-and-process", "GET", "/sp-period", GetDataSpPeriodNow},
		{"get-data-process", "GET", "/data-period", GetDataPeriod},
		{"add-quality-and-process", "POST", "/add/sp-period", AddDataSpPeriodNow},
		//billing 
		{"upload-file", "POST", "/billings/upload", UploadBilling},
		{"download-file", "GET", "/billings/download", DownloadBilling},
		//invoice
		{"upload-file", "POST", "/invoice/upload", UploadInvoice},
		//{"get-invoices", "GET", "/invoices", GetInvoicesByLogin},
		{"download-file", "GET", "/invoice/download", DownloadInvoice},
		//news attach
		{"upload-file", "POST", "/attachment/upload", UploadAttachments},
		//set comments analytic
		{"set-comment-file", "POST", "/set-comment", SetCommentFile},
		//approve sla
		{"set-comment-file", "GET", "/approve-sla", ApproveSla},
	}
}
