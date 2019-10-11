package fields

import "github.com/graphql-go/graphql"

func ID() *graphql.Field {
	return &graphql.Field{
		Type: graphql.ID,
	}
}

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
