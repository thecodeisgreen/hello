package args

import "github.com/graphql-go/graphql"

func Boolean() *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.Boolean,
	}
}

func ID() *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.ID,
	}
}

func String() *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.String,
	}
}
