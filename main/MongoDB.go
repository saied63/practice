package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {

}

var clint mongo.Client

//var mongoOptions options.ClientOptions

func CreateConnection() {
	if client == nil {
		mongoOptions := options.Client().ApplyURI("mongodb://saied63:111111@127.0.01:27017")
		timeout := time.Duration(10 * time.Second)
		mongoOptions.ServerSelectionTimeout = &timeout
		mongoOptions.SetMaxPoolSize(100)
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err = mongo.Connect(ctx, mongoOptions)

		defer cancel()
		defer clint.Disconnect(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func GetThisUser(dbname string, collectionName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return
	}
	collection := client.Database(dbname).Collection(collectionName)

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(ctx)
	var bsonDoc *bson.M
	for cursor.Next(ctx) {
		fmt.Println()
		err := cursor.Decode(&bsonDoc)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(bsonDoc)
		}
	}
}

/*
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
*/
