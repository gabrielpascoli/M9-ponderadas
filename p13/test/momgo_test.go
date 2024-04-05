package main

import (
	"context"
	"fmt"
	"os"

	config "iotsimkafka/config"

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
	err := bson.UnmarshalExtJSON([]
