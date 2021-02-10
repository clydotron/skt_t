package mongo

import (
	"clydotron/skt_t/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// CustomerModelX ...
type CustomerModelX struct {
	collection *mongo.Collection
}

// NewCustomerModelX ...
func NewCustomerModelX(c *mongo.Collection) *CustomerModelX {
	return &CustomerModelX{c}
}

// CreateCustomer ...
func (m *CustomerModelX) CreateCustomer(c models.Customer) error {

	c.ID = primitive.NewObjectID()

	_, err := m.collection.InsertOne(context.Background(), c)
	if err != nil {
		return err
	} //@
	return nil // can simplify?
}

// GetAllCustomers ...
func (m *CustomerModelX) GetAllCustomers() ([]*models.Customer, error) {

	cur, err := m.collection.Find(context.Background(), bson.M{})
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	defer cur.Close(context.Background())

	customers := []*models.Customer{}
	for cur.Next(context.Background()) {

		c := &models.Customer{}
		err := cur.Decode(c)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// GetCustomer ...
func (m *CustomerModelX) GetCustomer(id string) (*models.Customer, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//do something...
		return nil, err
	}

	c := &models.Customer{}
	err = m.collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// DeleteCustomer ...
func (m *CustomerModelX) DeleteCustomer(id string) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//do something...
		return err
	}

	_, err = m.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
