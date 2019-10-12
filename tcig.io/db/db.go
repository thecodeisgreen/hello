package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _dbHandler *mongo.Database

func Db() *mongo.Database {

	if _dbHandler == nil {
		fmt.Println("connection to ", os.Getenv("MONGO_DB_URL"))
		clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_DB_URL"))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.Background(), nil)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("connected")

		_dbHandler = client.Database(os.Getenv("MONGO_DATABASE"))
	}
	return _dbHandler
}
