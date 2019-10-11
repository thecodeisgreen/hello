package types

import (
	"hello/endpoint/mutations"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"signIn": mutations.SignInMutation,
			"signUp": mutations.SignUpMutation,
		},
	})
