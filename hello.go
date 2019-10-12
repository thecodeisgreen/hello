package main

import (
	"fmt"
	"hello/models/users"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hello/tcig.io/authentication"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	authentication.Init(router)

	user, err := users.GetOneByEmail("hubert.bettan@gmail.com")
	if err == users.ErrUserNotFound {
		fmt.Println("user not found")
	} else {
		fmt.Println("user found", user)
	}
	/*
		router.Use(middlewares.User())

		router.GET("/_/info", func(c *gin.Context) {
			sessionID, _ := c.Get("sessionID")
			c.JSON(200, gin.H{
				"version":   "1.0.0",
				"sessionID": sessionID,
			})
		})

		router.POST("/graphql", authentication.CheckAccess(), endpoint.GraphQLHandler())

		//router.NoRoute(hot_reloading.ReverseProxy())
	*/
	router.Run(os.Getenv("PORT"))

}
