package resolvers

import (
	"github.com/Tchayo/gql-tuts.git/internal/gql/models"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Resolver struct holds a connection to our db
type Resolver struct {
	DB *gorm.DB
}

// MessageResolver resolves our query through a db call to FindUserByID
func (r *Resolver) MessageResolver(p graphql.ResolveParams) (interface{}, error) {
	message := models.Message{}
	// Strip the name from the arguments and assert that it's a string
	id, ok := p.Args["id"].(uint32)
	if ok {
		messages, _ := message.FindMessageByID(r.DB, uint32(id))
		return messages, nil
	}
	return nil, nil

}

