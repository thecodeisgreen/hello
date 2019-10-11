package objects

import (
	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"signIn": SignInMutation,
			"signUp": SignUpMutation,
		},
	})
