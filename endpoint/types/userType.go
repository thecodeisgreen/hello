package types

import (
	"hello/endpoint/helper/fields"
	"hello/models/users"

	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: (func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(users.User).ID.Hex(), nil
				}),
			},
			"email": fields.String(),
		},
	},
)
