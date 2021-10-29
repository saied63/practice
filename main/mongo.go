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

var clientOption *options.ClientOptions
var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

func init() {
	ClientConnect()
}

type MongoPerson struct {
	Name string
	Age  int
	City string
}

func ClientConnect() bool {
	clientOption = options.Client().ApplyURI("mongodb://localhost:27017")
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var clientConnectionError error = nil
	client, clientConnectionError = mongo.Connect(ctx, clientOption)
	if clientConnectionError != nil {
		fmt.Printf("clinet connection  error ... ")
		return false
	}
	return true
}

func InsertOne(pack bson.D) bool {
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("firstDb").Collection("firstCollec")
	result, insertError := collection.InsertOne(ctx, pack)
	if insertError != nil {
		fmt.Println("error on inserting ... ")
		fmt.Println(insertError)
		return false
	} else {
		fmt.Printf("insert was seccussful and insert id is %d", result.InsertedID)
		return true
	}
}

func UpdateOne(pack bson.M, targetPack bson.M) int {
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("firstDb").Collection("firstCollec")
	updateResult, updateError := collection.UpdateOne(ctx, pack, targetPack)
	if updateResult != nil {
		fmt.Println("error on updating ... ")
		fmt.Println(updateError.Error())
		return 0
	} else {
		fmt.Printf("update was seccussful and insert id is %d \n", updateResult.UpsertedID)
		return int(updateResult.ModifiedCount)
	}
}

func SelectAll() {
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("firstDb").Collection("firstCollec")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("error to find collection ... ")
		return
	}
	defer cursor.Close(ctx)
	var result bson.D
	for cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(result)
	}
}

func FindeOne(one bson.M) bson.D {
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := client.Database("firstDb").Collection("firstCollec")
	singleResult := collection.FindOne(ctx, one)
	if singleResult.Err() != nil {
		fmt.Println("can't fine selected filter ... ")
		println(singleResult.Err())
		return nil
	}

	var result bson.D
	err := singleResult.Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(result)
	return result
}

func Ping() bool {
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pingErr := client.Ping(ctx, readpref.Primary())
	if pingErr != nil {
		fmt.Println("ping to mongo failed ... ")
		return false
	} else {
		fmt.Println("ping was seccussfull ... ")
		return true
	}
}

func Disconnect() {
	err := client.Disconnect(ctx)
	if err != nil {
		fmt.Println("context error ... ")
		fmt.Println(err)
	} else {
		fmt.Printf("client dose disconnect successful")
	}
}
