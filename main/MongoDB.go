package main

import (
	"context"
	"fmt"
	"time"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {

}

var clint mongo.Client;
var mongoOptions client.mongoOptions;

func CreateConnection(){
	if client==nil {
		mongoOptions = options.Client().ApplyURI("mongodb://saied63:111111@127.0.01:27017")
		mongoOptions.ServerSelectionTimeout = 10*time.Second;
		mongoOptions.SetMaxPooling(100);
		client, err = mongo.Connect(mongoOptions)
		defer clint.Disconnect()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func InsertOneDocument(databaseName string, collectionName string) {


	ctx, CancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer CancelFunc()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return;
	}
	db := client.Database(databaseName);
	collec:=db.Collection(collectionName);
	collec.
	ctx.InsertOne({
		name : "saied",
		family : "raj"
		car : "206"
	})
}
