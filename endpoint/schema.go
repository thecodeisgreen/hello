package endpoint

import (
	"hello/endpoint/mutations"
	"hello/endpoint/types"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    types.QueryType,
		Mutation: mutations.MutationType,
	},
)
