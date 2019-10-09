package types

import (
	"hello/endpoint/resolvers"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"saveUser": &graphql.Field{
				Type: SaveUserMutationResponseType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.ID,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userID, _ := p.Args["id"].(string)
					email, _ := p.Args["email"].(string)
					return resolvers.UserResolver().SaveOne(userID, email)
				},
			},
		},
	})
