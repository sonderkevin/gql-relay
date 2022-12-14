package gqlrelay

import (
	"context"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var nodeDefinitions *relay.NodeDefinitions
var userType *graphql.Object
var Schema graphql.Schema

func init() {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and it's type, return the object
			return GetUser(resolvedID.ID), nil
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			// based on the type of the value, return GraphQLObjectType
			return userType
		},
	})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "A user living in the world",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user.",
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	rootQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"users": &graphql.Field{
					Type: userType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return GetUsers(), nil
					},
				},
				"node": nodeDefinitions.NodeField,
			},
		},
	)

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
	if err != nil {
		panic(err)
	}
}
