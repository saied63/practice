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

var connectionString string

func init() {
	connectionString = "mongodb://saied63:111111@127.0.0.1:27017"
}

//var mongoOptions options.ClientOptions

func CreateConnection() mongo.Client {

	mongoOptions := options.Client().ApplyURI(connectionString)
	timeout := time.Duration(10 * time.Second)
	mongoOptions.ServerSelectionTimeout = &timeout
	mongoOptions.SetMaxPoolSize(100)
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if error := client.Disconnect(ctx); error != nil {
			panic(error)
		}
	}()
	//////////////
	client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	////////////////////////

	defer cancelCtx()
	//defer clint.Disconnect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return *client
}

func GetPing() bool {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	client := CreateConnection()
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return false
	} else {
		fmt.Println("ping goes right ..")
		return true
	}
}

func GetThisUser(dbname string, collectionName string) {
	mongoOptions := options.Client().ApplyURI(connectionString)
	timeout := time.Duration(10 * time.Second)
	mongoOptions.ServerSelectionTimeout = &timeout
	mongoOptions.SetMaxPoolSize(100)
	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelCtx()
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(dbname).Collection(collectionName)
	filter := bson.M{"_id": 7}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(ctx)
	var bsonDoc *bson.M
	// if cursor.Next(ctx) {
	// 	fmt.Println(cursor.RemainingBatchLength())
	// }
	for cursor.RemainingBatchLength() > 0 {
		fmt.Println(cursor.Next(ctx))
		err := cursor.Decode(&bsonDoc)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(bsonDoc)
		}
	}
	/*
		for cursor.Next(ctx) {
			fmt.Println()
			err := cursor.Decode(&bsonDoc)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(bsonDoc)
			}
		}
	*/
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
