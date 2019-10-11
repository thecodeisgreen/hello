package endpoint

import (
	"hello/endpoint/objects"

	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    objects.QueryType,
		Mutation: objects.MutationType,
	},
)
