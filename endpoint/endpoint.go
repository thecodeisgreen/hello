/*
 * subscription : https://medium.com/functional-foundry/building-graphql-servers-with-subscriptions-in-go-2a60f11dc9f5
 */

package endpoint

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func graphQLHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Init(router *gin.Engine) {
	group := router.Group("/graphql")
	group.POST("", graphQLHandler())
}
