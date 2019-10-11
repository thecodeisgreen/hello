package mutations

import (
	"hello/endpoint/helper/args"
	"hello/endpoint/helper/fields"
	"hello/endpoint/resolvers"

	"github.com/graphql-go/graphql"
)

var SignUpMutationResponseType *graphql.Object = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SignUpMutationResponse",
		Fields: graphql.Fields{
			"ok":    fields.Boolean(),
			"error": fields.String(),
		},
	},
)

var SignUpMutation *graphql.Field = &graphql.Field{
	Type: SignUpMutationResponseType,
	Args: graphql.FieldConfigArgument{
		"email":    args.String(),
		"password": args.String(),
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email, _ := p.Args["email"].(string)
		password, _ := p.Args["password"].(string)
		return resolvers.UserResolver().SignUp(email, password)
	},
}
