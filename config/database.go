package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func DBMongo(dbname string) *mongo.Database {
	connectionstr := os.Getenv("MONGOSTRCONNECT")
	if connectionstr == "" {
		panic(fmt.Errorf("MONGOSTRCONNECT ENV NOT FOUND"))
	}
	clay, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionstr))
	if err != nil {
		panic(fmt.Errorf("MongoConnect: %+v \n", err))
	}
	return clay.Database(dbname)

}
