package mongo

import (
	"clydotron/skt_t/pkg/models/controllers"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDBX ...
type MongoDBX struct {
	dbContext context.Context
	dbClient  *mongo.Client
	database  *mongo.Database
	//cc        *controllers.CustomerController
	Rentals *RentalModel
	RC      *controllers.RentalController

	Customers *CustomerModelX
	CC        *controllers.CustomerController
}

// ConnectToMongo ...
func (dbx *MongoDBX) ConnectToMongo() error {

	uri := os.Getenv("MONGODB_URI")
	// add check for this (doh)
	fmt.Println("connecting to mongo: ", uri)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	// do we want to ping to confirm?
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	dbx.dbContext = ctx
	dbx.dbClient = client
	dbx.database = client.Database("skt")

	fmt.Println("Connected to MongoDB")

	return nil
}

// InitControllers ...
func (dbx *MongoDBX) InitControllers() error {
	//dbx.cc = controllers.NewCustomerController(dbx.database.Collection("customers"))

	dbx.Rentals = NewRentalModel(dbx.database.Collection("rentals"))
	dbx.RC = controllers.NewRentalController(dbx.Rentals)

	dbx.Customers = NewCustomerModelX(dbx.database.Collection("customers"))
	dbx.CC = controllers.NewCustomerController(dbx.Customers)
	return nil
}
