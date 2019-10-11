package mutations

import (
	"hello/endpoint/helper/args"
	"hello/endpoint/helper/fields"
	"hello/endpoint/resolvers"

	"github.com/graphql-go/graphql"
)

var SignInMutationResponseType *graphql.Object = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SignInMutationResponse",
		Fields: graphql.Fields{
			"ok":    fields.Boolean(),
			"error": fields.String(),
		},
	},
)

var SignInMutation *graphql.Field = &graphql.Field{
	Type: SignInMutationResponseType,
	Args: graphql.FieldConfigArgument{
		"email":    args.String(),
		"password": args.String(),
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email, _ := p.Args["email"].(string)
		password, _ := p.Args["password"].(string)
		return resolvers.UserResolver().SignIn(email, password)
	},
}
