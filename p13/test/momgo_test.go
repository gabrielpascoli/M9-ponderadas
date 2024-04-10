package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	config "thefalloff/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() *mongo.Client {
	config.LoadEnv()

	mongodb := os.Getenv("MONGODB")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodb).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

type ApparentlySensor struct {
	Sensor string `json:"sensor"`
	Value  string `json:"value"`
}

func LightsPlease(data string) {
	client := ConnectMongoDB()

	defer client.Disconnect(context.TODO())

	var sensor ApparentlySensor

	// Access a MongoDB collection through a database
	err := bson.UnmarshalExtJSON([]byte(data), true, &sensor)
	if err != nil {
		panic(err)
	}

	collection := client.Database("Sensors").Collection(sensor.Sensor)

	// Insert a single document
	result, err := collection.InsertOne(context.TODO(), sensor)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", result.InsertedID)
}

func TestMongoDB(t *testing.T) {
	data := `{"sensor":"co2", "value":"10"}`
	LightsPlease(data)
}
