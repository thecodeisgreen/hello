package middlewares

import (
	"fmt"
	"hello/models/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sessionID string

		sessionCookieValue, err := c.Cookie("hello_session")
		fmt.Println(err)
		if err != nil {
			newSessionID, _ := uuid.NewRandom()
			sessionID = newSessionID.String()
			sessionCookieValue = "sessionId=" + sessionID + ";"
		}

		fmt.Println("-----")
		user, err := users.GetOneBySessionID(sessionID)
		fmt.Println(user)
		if user != nil {
			sessionCookieValue = sessionCookieValue + "email=" + user.Email + ";"
		}

		c.SetCookie(
			"hello_session",
			sessionCookieValue,
			36000,
			"/com.thecodeisgreen/hello",
			"localhost",
			false, // should be set to true when https is being used
			true,
		)

		c.Set("sessionID", sessionID)
		c.Next()
	}
}
