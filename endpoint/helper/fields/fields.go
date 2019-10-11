package fields

import "github.com/graphql-go/graphql"

func Boolean() *graphql.Field {
	return &graphql.Field{
		Type: graphql.Boolean,
	}
}

func String() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
	}
}
