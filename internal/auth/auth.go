package auth

import (
	"github.com/Tchayo/gql-tuts.git/internal/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Resolver struct holds a connection to our db
type DataB struct {
	DB *gorm.DB
}

type login struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// AuthenticatorFunction validate user
func (r *DataB) AuthenticatorFunction(c *gin.Context) (interface{}, error)  {
	user := models.Author{}

	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	email := loginVals.Email
	password := loginVals.Password

	output,err := user.Login(r.DB, email, password)
	if err != nil {
		return nil, err
		//return nil, jwt.ErrFailedAuthentication
	}
	return output, nil
}

// PayloadFunction map claims
func PayloadFunction(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.Author); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}

// IdentityHandlerFunction get request identity
func IdentityHandlerFunction(c *gin.Context) interface{}  {
	claims := jwt.ExtractClaims(c)
	return &models.Author{
		Email:    claims[jwt.IdentityKey].(string),
	}
}

// AuthorizatorFunction check authorizor
func AuthorizatorFunction(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.Author); ok && v.Email == "admin" {
		return true
	}
	return false
}

// UnauthorizedFunction returns unauthorized response
func UnauthorizedFunction(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}