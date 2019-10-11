/*
 * subscription : https://medium.com/functional-foundry/building-graphql-servers-with-subscriptions-in-go-2a60f11dc9f5
 */

package endpoint

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

/*
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
*/

func GraphQLHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		sessionID, _ := c.Get("sessionID")
		ctx := context.WithValue(c.Request.Context(), "sessionID", sessionID)
		h.ContextHandler(ctx, c.Writer, c.Request)
	}
}
