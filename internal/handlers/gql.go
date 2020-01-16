package handlers

import (
	"github.com/Tchayo/gql-tuts.git/internal/gql"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"net/http"
)

// Server will hold connection to the db
type Server struct {
	GqlSchema *graphql.Schema
}

// Binding from JSON
type RequestBody struct {
	Query string `json:"query" binding:"required"`
}

// GraphqlHandler defines the GQLGen GraphQL server handler
func (s *Server)GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	return func(c *gin.Context) {
		var json RequestBody
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := gql.ExecuteQuery(json.Query, *s.GqlSchema)
		c.JSON(http.StatusOK, result)
	}
}
