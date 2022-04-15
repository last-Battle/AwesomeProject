package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Student struct {
	Name string
	Age  int
}

func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}

func main() {

	s1 := Student{"小红", 12}
	s2 := Student{"小兰", 10}
	s3 := Student{"小黄", 11}

	uri := "mongodb://localhost:27017"
	dbName := "test"

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	c := client.Database(dbName).Collection("Student")
	// insert one
	insertResult, err := c.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted a single document:", insertResult.InsertedID)

	stus := []interface{}{s2, s3}
	insertManyResult, err := c.InsertMany(context.TODO(), stus)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents", insertManyResult.InsertedIDs)

	//update
	filter := bson.D{{"name", "小王子"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 19},
		}},
	}
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated %v documents ", updateResult.MatchedCount)

	// find
	var result Student
	err = c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("found a single document: %+v\n", result)

}
