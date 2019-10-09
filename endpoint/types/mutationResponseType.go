package types

import "github.com/graphql-go/graphql"

var SaveUserMutationResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SaveUserMutationRespone",
		Fields: graphql.Fields{
			"ok": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
