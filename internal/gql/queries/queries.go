package queries

import (
	"github.com/Tchayo/gql-tuts.git/internal/gql/resolvers"
	"github.com/Tchayo/gql-tuts.git/internal/gql/types"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *gorm.DB) *Root {
	// Create a resolver holding our database.
	resolver := resolvers.Resolver{DB: db}

	// Create a new Root that describes our base query set up
	// takes one argument called id
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "MessageQueries",
				Description: "Get message by ID",
				Fields: graphql.Fields{
					"message": &graphql.Field{
						// Slice of Message type found in types.go
						Type:    types.MessageType,
						Args:    graphql.FieldConfigArgument{
							"id": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
						},
						Resolve: resolver.MessageResolver,
					},
					"messages": &graphql.Field{
						// Slice of Message type found in types.go
						Type:    graphql.NewList(types.MessageType),
						Resolve: resolver.MessagesResolvers,
					},
				},
			},

		),
	}
	return &root
}
