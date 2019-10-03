package users

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hello/tcig.io/db"
)

// User struct representing user information
type User struct {
	Lastname  string
	Firstname string
}

var collection *mongo.Collection
var ctx context.Context

func init() {
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	collection = db.Db().Collection("users")
}

// Create insert a new user into users collection
func Create(
	Lastname string,
	Firstname string) {
	collection.InsertOne(ctx, User{
		Lastname:  Lastname,
		Firstname: Firstname})
}

// GetOneByID get user by _id
func GetOneByID(id string) User {
	var result User
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	collection.FindOne(ctx, filter).Decode(&result)
	return result
}

// Get query filter on users collection
func Get(filter bson.M) []User {
	var results []User

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)

	return results
}
