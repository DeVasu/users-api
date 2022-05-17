package tokens

import (
	"bookstore/token-jwt/domain/tokens"
	"bookstore/token-jwt/utils"
	"bookstore/token-jwt/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)








func GetUserByToken(c *gin.Context) {

	utils.EnableCors(c)
	// fmt.Println(c.Request.Header["Authorization"])

	t := tokens.Token{
		Token: c.Request.Header["Authorization"][0],
	}
	


	err := t.GetUserByToken()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	if t.IsExpired() {

		err := t.DeleteToken()
		if err != nil {
			c.JSON(err.Status, err)
			return
		}
		restErr := errors.NewBadRequestError("your token has expired, Please genereate a new one by logging in again")
		c.JSON(restErr.Status, restErr)
		return
	}

	t.Id = 0
	t.Token = ""
	t.Expires = 0

	c.JSON(http.StatusOK, t)

}