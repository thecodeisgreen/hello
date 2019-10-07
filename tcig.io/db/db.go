package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _dbHandler *mongo.Database

func Db() *mongo.Database {

	if _dbHandler == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		client, _ := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB_URL")))

		_dbHandler = client.Database(os.Getenv("MONGO_DATABASE"))
	}
	return _dbHandler
}
