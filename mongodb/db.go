package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Connection = "mongodb://0.0.0.0:27017"
)

func Getdb(dbname string, collectionname string) *mongo.Collection {
	// getting the mongo client
	client, err := mongo.NewClient(options.Client().ApplyURI(Connection))
	if err != nil {
		fmt.Println(err)
	}
	// waiting for 20 seconds for the connection to respond
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	// getting the collection from mongodb
	collection := client.Database(dbname).Collection(collectionname)
	return collection
}
