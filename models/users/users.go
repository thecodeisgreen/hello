package users

import (
	"context"
	"log"
	"reflect"
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

var _collection *mongo.Collection
var _ctx context.Context

func collection() *mongo.Collection {
	if _collection == nil {
		_ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		_collection = db.Db().Collection("users")
	}
	return _collection
}

func cursorToArray(cursor *mongo.Cursor) []User {
	var results []User
	for cursor.Next(_ctx) {
		var elem User
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(_ctx)

	return results
}

// Create insert a new user into users collection
func Create(
	Lastname string,
	Firstname string) {
	collection().InsertOne(_ctx, User{
		Lastname:  Lastname,
		Firstname: Firstname})
}

// GetOneByID get user by _id
func GetOneByID(id string) User {
	var result User
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectID}
	collection().FindOne(_ctx, filter).Decode(&result)
	return result
}

// Get query filter on users collection
func Get(filter bson.M) []User {

	cursor, err := collection().Find(_ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	return cursorToArray(cursor)
}

func Filter(users []User, filterFunc func(user User) bool) []User {

	var results []User

	for i := 0; i < len(users); i++ {
		if filterFunc(users[i]) {
			results = append(results, users[i])
		}
	}

	return results
}

func Map(users []User, mapFunc interface{}) interface{} {

	mapFuncValue := reflect.ValueOf(mapFunc)
	mapFuncType := mapFuncValue.Type()

	if mapFuncType.Kind() != reflect.Func || mapFuncType.NumIn() != 1 || mapFuncType.NumOut() != 1 {
		panic("second argument must be a map function")
	}

	if !mapFuncType.In(0).ConvertibleTo(reflect.TypeOf(User{})) {
		panic("Map function's argument is not compatible with type of array.")
	}

	resultSliceType := reflect.SliceOf(mapFuncType.Out(0))
	resultSlice := reflect.MakeSlice(resultSliceType, 0, len(users))

	for i := 0; i < len(users); i++ {
		resultSlice = reflect.Append(resultSlice, mapFuncValue.Call([]reflect.Value{reflect.ValueOf(users[i])})[0])
	}

	return resultSlice

}
