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

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/oauth2.v3/models"
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

func initClientStore() *store.ClientStore {
	clientStore := store.NewClientStore()

	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})

	clientStore.Set("Hubert", &models.Client{
		ID:     "Hubert",
		Secret: "",
		Domain: "http://localhost",
	})

	return clientStore
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
		manager := manage.NewDefaultManager()

		manager.MustTokenStorage(store.NewFileTokenStore(os.Getenv("ROOT_DIR") + "/_tokens/store"))
		//manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

		manager.MapClientStorage(initClientStore())

		// refresh_token wanted
		manager.SetClientTokenCfg(&manage.Config{
			AccessTokenExp:    time.Hour * 1,
			RefreshTokenExp:   time.Hour * 2,
			IsGenerateRefresh: true,
		})

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

		http.HandleFunc("/o/token", func(w http.ResponseWriter, r *http.Request) {
			srv.HandleTokenRequest(w, r)
		})

		http.HandleFunc("/o/refresh", func(w http.ResponseWriter, r *http.Request) {
			srv.HandleTokenRequest(w, r)
		})

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

	endpoint.Init(router)

	router.NoRoute(ReverseProxy())

	router.Run(os.Getenv("PORT"))

}
