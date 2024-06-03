package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Restaurant struct {
	Name         string
	RestaurantId string        `bson:"restaurant_id,omitempty"`
	Cuisine      string        `bson:"cuisine,omitempty"`
	Address      interface{}   `bson:"address,omitempty"`
	Borough      string        `bson:"borough,omitempty"`
	Grades       []interface{} `bson:"grades,omitempty"`
}

func main() {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Println("MongoDB URL not working")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Could not disconnect from DB")
		}
	}()
	coll := client.Database("sample_restaurants").Collection("restaurants")
	newRestaurant := Restaurant{Name: "8282", Cuisine: "Korean"}
	_, err = coll.InsertOne(context.TODO(), newRestaurant)
	if err != nil {
		log.Fatalf(err.Error())
	}
	filter := bson.D{{"name", "8282"}}
	// Retrieves the first matching document
	var result Restaurant
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("Document inserted with ID: %v\n", result)
}
