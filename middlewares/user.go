package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getParams(c *gin.Context) url.Values {
	var params url.Values = url.Values{}

	value, err := c.Cookie("HELLO_SESSION")
	if err == http.ErrNoCookie {
		newSessionID, _ := uuid.NewRandom()
		params.Add("sessionID", newSessionID.String())
	} else {
		decodedValue, err := url.QueryUnescape(value)
		if err != nil {
			log.Fatal(err)
		}
		params, err = url.ParseQuery(decodedValue)
		fmt.Println(params)
	}

	return params
}

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := getParams(c)

		/*
			user, _ := users.GetOneBySessionID(params.Get("sessionID"))

			fmt.Println(user)
			if user != nil {
				params.Add("email", user.Email)
			}
		*/

		c.SetCookie(
			"HELLO_SESSION",
			params.Encode(),
			36000,
			"",
			"",
			false, // should be set to true when https is being used
			false,
		)

		c.Set("sessionID", params.Get("sessionID"))
		c.Next()
	}
}
