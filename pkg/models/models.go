package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ContactInfo ...
// we can get fancier eventually, right now just need name/phone
type ContactInfo struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

// Keg ...
type Keg struct {
	ID   primitive.ObjectID
	Code string
	Size int
}

// KegRental ...
type KegRental struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID primitive.ObjectID `json:"customerID"`
	KegID      primitive.ObjectID `json:"kegID"`
	TimeStamp  time.Time          `json:"timeStamp"`
	Contents   string             `json:"contents"`
}

// Customer ...
type Customer struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string
	Email string
	//Info ContactInfo
}

// RentalModel ...
type RentalModel interface {
	CreateRental(r KegRental) error
	GetRental(id string) (*KegRental, error)
	GetAllRentals() ([]*KegRental, error)
	DeleteRental(id string) error
}

// CustomerModel ...
type CustomerModel interface {
	CreateCustomer(r Customer) error
	GetCustomer(id string) (*Customer, error)
	GetAllCustomers() ([]*Customer, error)
	DeleteCustomer(id string) error
}
