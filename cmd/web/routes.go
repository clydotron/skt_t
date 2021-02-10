package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) createRouter() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//	v1 := router.Group("/api/v1")

	return mux
}

func (app *application) createRouterP() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMW := alice.New(app.session.Enable)
	//dynamicMW := alice.New()

	// with pat order matters
	mux := pat.New()
	mux.Get("/", dynamicMW.ThenFunc(app.home))
	mux.Get("/rentals/create", dynamicMW.ThenFunc(app.createRentalForm))
	mux.Post("/rentals/create", dynamicMW.ThenFunc(app.createRental))
	mux.Get("/rentals", dynamicMW.ThenFunc(app.showAllRentals))
	mux.Get("/rentals/return/:id", dynamicMW.ThenFunc(app.handleRentalReturn))
	mux.Get("/rentals/:id", dynamicMW.ThenFunc(app.showRental))

	mux.Get("/customers", dynamicMW.ThenFunc(app.showAllCustomers))
	mux.Get("/customers/create", dynamicMW.ThenFunc(app.createCustomerForm))
	mux.Post("/customers/create", dynamicMW.ThenFunc(app.createCustomer))
	mux.Get("/customers/:id", dynamicMW.ThenFunc(app.showCustomer))
	mux.Get("/customers/:id/delete", dynamicMW.ThenFunc(app.deleteCustomer))
	mux.Get("/customers/:id/purchase", dynamicMW.ThenFunc(app.customerPurchase))

	mux.Get("/api/v1/rentals", http.HandlerFunc(app.dbx.RC.GetAllRentals))
	mux.Post("/api/v1/rentals", http.HandlerFunc(app.dbx.RC.CreateRental))

	mux.Get("/api/v1/customer", http.HandlerFunc(app.dbx.CC.GetAllCustomers))
	mux.Post("/api/v1/customer", http.HandlerFunc(app.dbx.CC.CreateCustomer))
	mux.Del("/api/v1/customer/:id", http.HandlerFunc(app.dbx.CC.DeleteCustomer))

	//mux.Get("/api/v1/customers", dynamicMW.ThenFunc(app.dbx.CC.getAllCustomers))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
