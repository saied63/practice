package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	fmt.Println("mongo db from scratch ... ")

}

func MackeConnectionToMOnog() bool {
	mongoClientOption := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	var err error
	mongoClient, err = mongo.Connect(ctx, mongoClientOption)
	if err != nil {
		fmt.Printf("error on connecting to mongo %s", err.Error())
		return false
	}
	return true
}

func PingToMongo() bool {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	mongoClient.Ping(ctx, readpref.Primary())
	return false
}
