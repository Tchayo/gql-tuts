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
				Name: "Query",
				Fields: graphql.Fields{
					"messages": &graphql.Field{
						// Slice of Message type found in types.go
						Type:    graphql.NewList(types.MessageType),
						Args:    nil,
						Resolve: resolver.MessageResolver,
					},
				},
			},

		),
	}
	return &root
}
