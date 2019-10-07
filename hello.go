package main

import (
	"fmt"
	"hello/models/users"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func filterFunc(user users.User) bool {
	return user.Firstname == "augustine"
}

type Friend struct {
	Lastname          string
	Firstname         string
	PreviousFirstname string
}

func mapFunc(user users.User) Friend {
	return Friend{
		Lastname:          user.Lastname,
		Firstname:         "Lilou",
		PreviousFirstname: user.Firstname}
}

func mapUsers(users []users.User) []Friend {
	var results []Friend

	for i := 0; i < len(users); i++ {
		user := users[i]
		results = append(results, Friend{
			Lastname:          user.Lastname,
			Firstname:         "Lilou",
			PreviousFirstname: user.Firstname})
	}

	return results
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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	manager := manage.NewDefaultManager()

	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))

}