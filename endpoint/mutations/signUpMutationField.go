package mutations

import (
	"hello/endpoint/helper/args"
	"hello/endpoint/helper/fields"
	"hello/endpoint/resolvers"
	"hello/endpoint/types"

	"github.com/graphql-go/graphql"
)

var SignUpMutationResponseType *graphql.Object = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "SignUpMutationResponse",
		Fields: graphql.Fields{
			"ok":    fields.Boolean(),
			"error": fields.String(),
			"user": &graphql.Field{
				Type: types.UserType,
			},
		},
	},
)

var SignUpMutationField *graphql.Field = &graphql.Field{
	Type: SignUpMutationResponseType,
	Args: graphql.FieldConfigArgument{
		"email":    args.String(),
		"password": args.String(),
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email, _ := p.Args["email"].(string)
		password, _ := p.Args["password"].(string)
		return resolvers.UserResolver(p.Context).SignUp(email, password)
	},
}
