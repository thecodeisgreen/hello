package objects

import (
	"hello/endpoint/helper/args"
	"hello/endpoint/resolvers"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"id": args.ID(),
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return resolvers.UserResolver(p.Context).User(), nil
				},
			},
		},
	})
