package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"newBackend/config"
	"newBackend/controller"
)

var Router = mux.NewRouter()

func makeHandler(method, endpoint string, fn func(http.ResponseWriter, *http.Request)) {
	Router.HandleFunc(fmt.Sprintf("/%v/%v", config.AppName, endpoint), fn).Methods(method)
}

func InitHandlers() {
	productHandlers()
	shopHandlers()
	billHandler()
	paymentHandler()
	supplierHandlers()
	customersHandlers()
	staffHandlers()
}

func paymentHandler() {
	makeHandler(http.MethodGet, "record_payment",
		controller.RecordPayment)
}

func shopHandlers() {
	makeHandler(http.MethodGet, "get_shops",
		controller.GetShops)
}

func billHandler() {
	makeHandler(http.MethodGet, "make_bill",
		controller.HandleSaveBill)
	makeHandler(http.MethodGet, "get_bills",
		controller.HandleGetBills)
}

func productHandlers() {
	makeHandler(http.MethodGet, "get_products",
		controller.GetProducts)
	makeHandler(http.MethodGet, "update_products",
		controller.UpdateProducts)
	makeHandler(http.MethodGet, "add_products",
		controller.AddProducts)
	makeHandler(http.MethodGet, "delete_products",
		controller.DeleteProducts)
}

func supplierHandlers() {
	makeHandler(http.MethodGet, "get_suppliers",
		controller.GetSuppliers)
	makeHandler(http.MethodGet, "delete_suppliers",
		controller.DeleteSuppliers)
}

func customersHandlers() {
	makeHandler(http.MethodGet, "get_customers",
		controller.GetCustomers)
}

func staffHandlers() {
	makeHandler(http.MethodGet, "get_staff",
		controller.GetStaff)
}
