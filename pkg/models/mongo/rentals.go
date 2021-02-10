package mongo

import (
	"clydotron/skt_t/pkg/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// RentalModel ...
type RentalModel struct {
	collection *mongo.Collection
}

// NewRentalModel ...
func NewRentalModel(c *mongo.Collection) *RentalModel {
	return &RentalModel{c}
}

// CreateRental ...
func (m *RentalModel) CreateRental(r models.KegRental) error {

	r.ID = primitive.NewObjectID()
	r.TimeStamp = time.Now()

	_, err := m.collection.InsertOne(context.Background(), r)
	if err != nil {
		return err
	} //@
	return nil // can simplify?
}

// GetRental ...
func (m *RentalModel) GetRental(id string) (*models.KegRental, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//do something...
		return nil, err
	}

	rental := &models.KegRental{}

	filter := bson.M{"_id": objID}
	err = m.collection.FindOne(context.Background(), filter).Decode(rental)
	if err != nil {
		return nil, err
	}
	return rental, nil
}

// GetAllRentals ...
func (m *RentalModel) GetAllRentals() ([]*models.KegRental, error) {

	cur, err := m.collection.Find(context.Background(), bson.M{})
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	rentals := []*models.KegRental{}
	for cur.Next(context.Background()) {

		r := &models.KegRental{}
		err := cur.Decode(r)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, r)
	}
	return rentals, nil
}

// GetAllRentalsByCustomer ...
func (m *RentalModel) GetAllRentalsByCustomer(id string) ([]*models.KegRental, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"customerid": objID}
	cur, err := m.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	rentals := []*models.KegRental{}
	for cur.Next(context.Background()) {

		r := &models.KegRental{}
		err := cur.Decode(r)
		if err != nil {
			return nil, err
		}
		rentals = append(rentals, r)
	}
	return rentals, nil
}

// DeleteRental ...
func (m *RentalModel) DeleteRental(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//do something...
		return err
	}

	//filter := bson.M{"_id": objID}

	_, err = m.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		//c.String(http.StatusInternalServerError, "Keg: DeleteOne failed:")
		return err
	}
	return nil
}
