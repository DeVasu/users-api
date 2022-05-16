package users

import (
	"bookstore/token-jwt/domain/tokens"
	"bookstore/token-jwt/domain/users"
	"bookstore/token-jwt/utils/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	saveErr := user.Save()
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}


func Login(c *gin.Context) {
	var request users.User

	if err := c.BindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	err := request.FindByEmailAndPassword()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	resToken := tokens.Token{
		UserId: request.Id,
	}
	err = resToken.GetByUserId()
	if err != nil {

		if strings.Contains(err.Message, "token does not exists for this user") {
			err = resToken.Generate(request.Id)
			if err != nil {
				c.JSON(err.Status, err)
				return
			}
			resToken.Id = 0
			resToken.UserId = 0
			c.JSON(http.StatusOK, resToken)
			return
		}

		c.JSON(err.Status, err)
		return
	}
	resToken.Id = 0
	resToken.UserId = 0
	c.JSON(http.StatusOK, resToken)

}