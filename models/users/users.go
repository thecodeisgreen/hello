package users

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hello/tcig.io/db"
)

// User struct representing user information

type NewUser struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	SessionID string             `bson:"sessionID"`
	ForceID   *string
}

var ErrUserNotFound error = errors.New("user: user not found")

var _collection *mongo.Collection

func getContext() context.Context {
	return context.TODO()
}
func collection() *mongo.Collection {
	if _collection == nil {
		_collection = db.Db().Collection("users")
	}
	return _collection
}

func cursorToArray(cursor *mongo.Cursor) []User {
	var results []User
	for cursor.Next(getContext()) {
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
	cursor.Close(getContext())

	return results
}

// Create insert a new user into users collection
func CreateOne(newUser NewUser) *User {
	res, err := collection().InsertOne(getContext(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	user, _ := GetOneByID(res.InsertedID.(primitive.ObjectID).Hex())

	return user
}

// GetOneByID get user by _id
func GetOneByID(id string) (*User, error) {
	return getOne(bson.M{"_id": id})
}

// GetOneByID get user by _id
func GetOneByEmail(email string) (*User, error) {
	return getOne(bson.M{"email": email})
}

// GetOneBySessionID
func GetOneBySessionID(sessionID string) (*User, error) {
	return getOne(bson.M{"sessionID": sessionID})
}

func getOne(filter bson.M) (*User, error) {
	var result User
	fmt.Println("---->", filter)
	err := collection().FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("---->")
	if err == mongo.ErrNoDocuments || result.ID == primitive.NilObjectID {
		return nil, ErrUserNotFound
	}
	return &result, nil
}

// Get query filter on users collection
func Get(filter bson.M) []User {

	cursor, err := collection().Find(getContext(), filter)
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

func (user *User) SetSessionID(sessionID string) {
	user.SessionID = sessionID
}

func (user *User) Save() {
	collection().UpdateOne(
		getContext(),
		bson.M{"_id": user.ID},
		user,
	)
}
