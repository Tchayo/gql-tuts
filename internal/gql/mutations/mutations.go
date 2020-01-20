package mutations

import (
	"github.com/Tchayo/gql-tuts.git/internal/gql/resolvers"
	"github.com/Tchayo/gql-tuts.git/internal/gql/types"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

type Root struct {
	Mutation graphql.Object
}

// NewRootMutation returns message creation mutation
func NewRootMutation(db *gorm.DB) *graphql.Object {
	// Create a resolver holding our database.
	resolver := resolvers.Resolver{DB: db}

	// Create a new Root that describes our base query set up
	rootMutation := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "RootMutation",
			Fields: graphql.Fields{
				"newUser": &graphql.Field{
					Type: types.NewAuthorType,
					Args: graphql.FieldConfigArgument{
						"username": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"email": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"password": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve:     resolver.CreateUserResolver,
					Description: "Create new system user",
				},
				"createMessage": &graphql.Field{
					Type: types.MessageType,
					Args: graphql.FieldConfigArgument{
						"shortcode": &graphql.ArgumentConfig{
							Type:        graphql.NewNonNull(graphql.String),
							Description: "",
						},
						"number": &graphql.ArgumentConfig{
							Type:        graphql.NewNonNull(graphql.String),
							Description: "",
						},
						"message": &graphql.ArgumentConfig{
							Type:        graphql.NewNonNull(graphql.String),
							Description: "",
						},
						"scheduled": &graphql.ArgumentConfig{
							Type:         graphql.Boolean,
							DefaultValue: false,
							Description:  "",
						},
						"schedule_time": &graphql.ArgumentConfig{
							Type:         graphql.DateTime,
							DefaultValue: nil,
							Description:  "",
						},
					},
					Resolve:     resolver.CreateMessageResolver,
					Description: "Validate and create a new message",
				},
			},
			Description: "Create message Mutation",
		},
	)

	return rootMutation
}
