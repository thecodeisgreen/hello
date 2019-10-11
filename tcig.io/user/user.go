package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func User() gin.HandlerFunc {
	return func(c *gin.Context) {
		var sessionID string
		sessionID, err := c.Cookie("sessionID")
		fmt.Println(err)
		if err != nil {
			newSessionID, _ := uuid.NewRandom()
			sessionID = newSessionID.String()
			c.SetCookie(
				"sessionID",
				sessionID,
				36000,
				"/com.thecodeisgreen/hello",
				"localhost",
				false, // should be set to true when https is being used
				true,
			)
		}

		c.Set("sessionID", sessionID)
		c.Next()
	}
}
