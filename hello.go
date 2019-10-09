package main

import (
	"fmt"
	"hello/endpoint"
	"hello/models/users"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"

	"hello/tcig.io/authentication"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func ReverseProxy() gin.HandlerFunc {

	target := "localhost:3000"

	return func(c *gin.Context) {
		fmt.Println(c.Request)
		director := func(req *http.Request) {
			//req = c.Request
			//req = r
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Host = target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/*
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			token, err := srv.ValidationBearerToken(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Write([]byte(token.GetScope() + " Doing good?"))
		})
	*/

	router := gin.Default()

	router.GET("/firstname", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"firstname": "hubert",
		})
	})

	router.GET("/lastname", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"lastname": "bettan",
		})
	})

	router.GET("/version", func(c *gin.Context) {
		authentication.CheckAccess(c, func(scope string) {
			c.JSON(200, gin.H{
				"version": "1.0.0",
				"scope":   scope,
			})
		})
	})

	authentication.Init(router)
	endpoint.Init(router)

	router.NoRoute(ReverseProxy())

	router.Run(os.Getenv("PORT"))

}
