package controllers

import (
	"clydotron/skt_t/pkg/models"
	"encoding/json"
	"net/http"
)

// CustomerController x
type CustomerController struct {
	m models.CustomerModel
}

// NewCustomerController create new customer controller
func NewCustomerController(m models.CustomerModel) *CustomerController {
	return &CustomerController{m}
}

// CreateCustomer ...
func (cc *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	c := models.Customer{}
	json.NewDecoder(r.Body).Decode(&c)

	err := cc.m.CreateCustomer(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//@todo, return the new json object?
	//fmt.Println("K:", k)
	w.Write([]byte("Customer created"))
}

// GetAllCustomers ...
func (cc *CustomerController) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := cc.m.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// DeleteCustomer ...
func (cc *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	err := cc.m.DeleteCustomer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Customer deleted"))
}
