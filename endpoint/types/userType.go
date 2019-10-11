package types

import (
	"hello/endpoint/helper/fields"

	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    fields.String(),
			"email": fields.String(),
		},
	},
)
