package types

import (
	"hello/endpoint/helper/fields"
	"hello/models/users"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/graphql-go/graphql"
)

type UserTypeStruct struct {
	ID    *string
	Email string `bson:"email"`
}

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
				Resolve: (func(p graphql.ResolveParams) (interface{}, error) {
					ID := p.Source.(users.User).ID
					if ID == primitive.NilObjectID {
						return nil, nil
					}
					return ID.Hex(), nil
				}),
			},
			"email": fields.String(),
		},
	},
)
