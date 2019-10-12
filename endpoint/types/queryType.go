package types

import (
	"hello/endpoint/helper/args"
	"hello/endpoint/resolvers"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"UserSession": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"id": args.ID(),
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					user := resolvers.UserResolver(p.Context).User()
					return user, nil
				},
			},
		},
	})
