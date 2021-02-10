package controllers

import (
	"clydotron/skt_t/pkg/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// RentalController x
type RentalController struct {
	collection *mongo.Collection
}

// NewRentalController create new rental controller
func NewRentalController(c *mongo.Collection) *RentalController {
	return &RentalController{c}
}

// CreateRental ...
func (rc *RentalController) CreateRental(c *gin.Context) {

	r := models.KegRental{}
	json.NewDecoder(c.Request.Body).Decode(&r)

	r.ID = primitive.NewObjectID()

	_, err := rc.collection.InsertOne(context.Background(), r)
	if err != nil {
		log.Fatal(err)
		return
	} //@todo use c.String

	// should we use the customer from the response?
	c.JSON(http.StatusCreated, r)
}

// GetAllRentals ...
func (rc *RentalController) GetAllRentals(c *gin.Context) {

	cur, err := rc.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	rentals := []models.KegRental{}

	for cur.Next(context.Background()) {

		r := models.KegRental{}
		err := cur.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}
		rentals = append(rentals, r)
	}

	c.JSON(http.StatusOK, rentals)
}

func (rc *RentalController) GetAllRentalsG(c *gin.Context) {

	cur, err := rc.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	rentals := []models.KegRental{}

	for cur.Next(context.Background()) {

		r := models.KegRental{}
		err := cur.Decode(&r)
		if err != nil {
			log.Fatal(err)
		}
		rentals = append(rentals, r)
	}

	c.JSON(http.StatusOK, rentals)
}

//get
// GetRental ...
func (rc *RentalController) GetRental(c *gin.Context) {

}

//update
func (rc *RentalController) UpdateRental(c *gin.Context) {

}

//delete
func (rc *RentalController) DeleteRental(c *gin.Context) {

}
