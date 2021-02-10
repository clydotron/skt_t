package main

import (
	"clydotron/skt_t/pkg/forms"
	"clydotron/skt_t/pkg/models"
	"fmt"
	"html/template"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) createRentalForm(w http.ResponseWriter, r *http.Request) {
}

func (app *application) createRental(w http.ResponseWriter, r *http.Request) {
}

func (app *application) showRental(w http.ResponseWriter, r *http.Request) {
}

func (app *application) showAllRentals(w http.ResponseWriter, r *http.Request) {

	rentals, err := app.dbx.Rentals.GetAllRentals()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "rentals.page.tmpl", &templateData{
		Rentals: rentals,
	})
}

func (app *application) createCustomerForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "customer_create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createCustomer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content.
	form := forms.New(r.PostForm)
	form.Required("name", "email")
	form.MaxLength("name", 100)

	name := form.Get("name")
	email := form.Get("email")

	// If the form isn't valid, redisplay the template passing in the
	// form.Form object as the data.
	if !form.Valid() {
		app.render(w, r, "customer_create.page.tmpl", &templateData{Form: form})
		return
	}

	c := models.Customer{
		Name:  name,
		Email: email,
	}

	// Create a new snippet record in the database using the form data.
	err = app.dbx.Customers.CreateCustomer(c)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Customer successfully created!")

	//@todo need a way to get the ID back from create
	//http.Redirect(w, r, fmt.Sprintf("/customer/%d", id), http.StatusSeeOther)
	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}

func (app *application) showAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := app.dbx.Customers.GetAllCustomers()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper.
	app.render(w, r, "customers.page.tmpl", &templateData{
		Customers: customers,
	})
}

func (app *application) showCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	customer, err := app.dbx.Customers.GetCustomer(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	rentals, err := app.dbx.Rentals.GetAllRentalsByCustomer(id)
	if err != nil {
		//do something...
		fmt.Println("failed:", err)
	} else {
		fmt.Println("#Rentals:", len(rentals))
	}

	app.render(w, r, "customer.page.tmpl", &templateData{
		Customer: customer,
		Rentals:  rentals,
	})
}

func (app *application) deleteCustomer(w http.ResponseWriter, r *http.Request) {

	err := app.dbx.Customers.DeleteCustomer(r.URL.Query().Get(":id"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	// @todo we dont know their name at this point...
	app.session.Put(r, "flash", "Customer <name> deleted.")

	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}

// not crazy about the name..
func (app *application) customerPurchase(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	cObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}
	// show a list of kegs?

	// @todo right now jsut create a rental object (without the keg info)
	kr := models.KegRental{
		CustomerID: cObjID,
		Contents:   "hoppy goodness",
	}
	app.session.Put(r, "flash", "Rental created.")

	app.dbx.Rentals.CreateRental(kr)

	http.Redirect(w, r, "/rentals", http.StatusSeeOther)
}

func (app *application) handleRentalReturn(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")

	err := app.dbx.Rentals.DeleteRental(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//w.Write([]byte("returned"))

	//show something?
	//flash?
	app.session.Put(r, "flash", "Rental returned")

	// navigate somewhere else?
	// Redirect the user to the rentals page.
	http.Redirect(w, r, "/rentals", http.StatusSeeOther)
}
