package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Example struct {
	Type      string    `json:"type" bson:"type"`
	CreatedAT time.Time `json:"created_at" bson:"created_at"`
	Created   int64     `json:"created" bson:"created"`
}

func main() {
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("test").Collection("coll")
	//coll.InsertOne(context.TODO(), &Example{
	//	Type:      "B",
	//	CreatedAT: time.Now(),
	//	Created:   time.Now().UnixNano(),
	//})
	//coll.InsertOne(context.TODO(), &Example{
	//	Type:      "C",
	//	CreatedAT: time.Now(),
	//	Created:   time.Now().UnixNano(),
	//})
	//coll.InsertOne(context.TODO(), &Example{
	//	Type:      "B",
	//	CreatedAT: time.Now(),
	//	Created:   time.Now().UnixNano(),
	//})
	//coll.InsertOne(context.TODO(), &Example{
	//	Type:      "C",
	//	CreatedAT: time.Now(),
	//	Created:   time.Now().UnixNano(),
	//})
	//coll.InsertOne(context.TODO(), &Example{
	//	Type:      "B",
	//	CreatedAT: time.Now(),
	//	Created:   time.Now().UnixNano(),
	//})

	cursor, err := coll.Find(context.TODO(), bson.M{}, options.Find().SetSort(
		bson.M{"type": 1, "created": -1}))
	if err != nil {
		panic(err)
	}

	var examples []*Example
	if err := cursor.All(ctx, &examples); err != nil {
		panic(err)
	}

	for _, e := range examples {
		fmt.Printf("%v\n", e)
	}
}
