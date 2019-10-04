package main

import (
	"fmt"
	"reflect"
	"time"
	"hello/models/users"

	"go.mongodb.org/mongo-driver/bson"
)

func filterFunc (user users.User) bool {
	return user.Firstname == "augustine"
}

type Friend struct {
	Lastname string
	Firstname string
	PreviousFirstname string
}

func mapFunc (user users.User) Friend {
	return Friend{
		Lastname: user.Lastname, 
		Firstname: "Lilou",
		PreviousFirstname: user.Firstname}
}

func mapUsers(users []users.User) []Friend {
	var results []Friend

	for i:=0; i<len(users); i++ {
		user := users[i]
		results = append(results, Friend{
			Lastname: user.Lastname, 
			Firstname: "Lilou",
			PreviousFirstname: user.Firstname})
	}

	return results
}

func main() {
	var start time.Time
	//users.Create("bettan", "hubert")
	//users.Create("bettan", "augustine")
	allUsers := users.Get(bson.M{"lastname": "bettan"})
	start = time.Now()
	fmt.Println(filter(allUsers, filterFunc))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(users.Filter(allUsers, filterFunc))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(users.Map(allUsers, mapFunc))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(mapUsers(allUsers))
	fmt.Println(time.Since(start))

}

func isBool(value interface{}) bool {
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Bool
}

func filter(items interface{}, filterFunc interface{}) interface{} {
	filterFuncValue := reflect.ValueOf(filterFunc)
	itemsValue := reflect.ValueOf(items)

	itemsType := itemsValue.Type()
	itemsElemType := itemsType.Elem()

	filterFuncType := filterFuncValue.Type()

	if filterFuncType.Kind() != reflect.Func || filterFuncType.NumIn() != 1 || filterFuncType.NumOut() != 1 {
		panic("second argument must be a filter function")
	}

	if !itemsElemType.ConvertibleTo(filterFuncType.In(0)) {
		panic("Map function's argument is not compatible with type of array.")
	}

	resutSliceType := reflect.SliceOf(itemsElemType)
	resultSlice := reflect.MakeSlice(resutSliceType, 0, itemsValue.Len())

	for i := 0; i < itemsValue.Len(); i++ {
		keep := filterFuncValue.Call([]reflect.Value{itemsValue.Index(i)})[0]

		if keep.Bool() {
			resultSlice = reflect.Append(resultSlice, itemsValue.Index(i))
		}
	}

	return resultSlice.Interface()
}
