package main

import (
	"hello/endpoint"
	"log"
	"os"

	"hello/tcig.io/authentication"
	"hello/tcig.io/hot_reloading"
	"hello/tcig.io/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	authentication.Init(router)

	router.Use(user.User())

	router.GET("/_/info", func(c *gin.Context) {
		sessionID, _ := c.Get("sessionID")
		c.JSON(200, gin.H{
			"version":   "1.0.0",
			"sessionID": sessionID,
		})
	})

	router.POST("/graphql", authentication.CheckAccess(), endpoint.GraphQLHandler())

	router.NoRoute(hot_reloading.ReverseProxy())

	router.Run(os.Getenv("PORT"))

}
