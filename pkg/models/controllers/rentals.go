package controllers

import (
	"clydotron/skt_t/pkg/models"
	"encoding/json"
	"net/http"
)

// @todo the model needs to be an interface!
type RentalController struct {
	m models.RentalModel
}

// NewCustomerController create new customer controller
func NewRentalController(m models.RentalModel) *RentalController {
	return &RentalController{m}
}

// CreateRental ...
func (rc *RentalController) CreateRental(w http.ResponseWriter, r *http.Request) {

	k := models.KegRental{}
	json.NewDecoder(r.Body).Decode(&k)

	err := rc.m.CreateRental(k)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//@todo, return the new json object? (should we get it from the DB?)
	//fmt.Println("K:", k)
	w.Write([]byte("Rental created"))
}

// Get ...
func (rc *RentalController) Get(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get(":id")
	//fmt.Println("Get:", id)

	rental, err := rc.m.GetRental(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(rental)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetAllRentals ...
func (rc *RentalController) GetAllRentals(w http.ResponseWriter, r *http.Request) {

	rentals, err := rc.m.GetAllRentals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(rentals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
