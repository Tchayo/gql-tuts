package resolvers

import (
	"errors"
	"github.com/Tchayo/gql-tuts.git/internal/models"
	"github.com/Tchayo/gql-tuts.git/internal/utils"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"time"
)

// Resolver struct holds a connection to our db
type Resolver struct {
	DB *gorm.DB
}

// CreateUserResolver create new system user
func (r *Resolver) CreateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	// marshal and cast argument values
	username, _ := p.Args["username"].(string)
	email, _ := p.Args["email"].(string)
	password, _ := p.Args["password"].(string)

	if username == "" || email == "" || password == "" {
		return nil, errors.New("all values required")
	}

	err := utils.ValidateFormat(email)
	if err != nil {
		return nil, err
	}

	if len(password) < 6 {
		return nil, errors.New("password too short")
	}

	newUser := models.Author{
		Name:     username,
		Email:    email,
		Password: password,
	}

	output, err := newUser.SaveUser(r.DB)
	if err != nil {
		return nil, err
	}
	return output, nil

}

// MessageResolver resolves our query through a db call to FindUserByID
func (r *Resolver) MessageResolver(p graphql.ResolveParams) (interface{}, error) {
	message := models.Message{}
	// Strip the name from the arguments and assert that it's a string
	id, ok := p.Args["id"].(int)
	if ok {
		messages, _ := message.FindMessageByID(r.DB, uint32(id))
		return messages, nil
	}
	return nil, nil

}

func (r *Resolver) MessagesResolvers(p graphql.ResolveParams) (interface{}, error) {
	message := models.Message{}
	messages, err := message.FindAllMessages(r.DB)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *Resolver) CreateMessageResolver(p graphql.ResolveParams) (interface{}, error) {
	// marshal and cast argument values
	short, _ := p.Args["shortcode"].(string)
	number, _ := p.Args["number"].(string)
	text, _ := p.Args["message"].(string)
	scheduled, _ := p.Args["message"].(bool)
	sTime, _ := p.Args["schedule_time"].(time.Time)

	if short == "" || number == "" || text == "" {
		return nil, errors.New("shortcode, number and message required")
	}

	// perform mutation
	newMessage := models.Message{
		ID:           0,
		ShortCode:    short,
		Number:       number,
		Message:      text,
		Scheduled:    scheduled,
		ScheduleTime: sTime,
		CreatedAt:    time.Time{},
	}

	output, err := newMessage.SaveMessage(r.DB)
	if err != nil {
		return nil, err
	}
	return output, nil
}
