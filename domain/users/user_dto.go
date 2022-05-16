package users

// data transfer object

import (
	"bookstore/token-jwt/utils/errors"
	"strings"
)


type Users []User

type User struct {
	Id          int64  `json:"id"`
	Name   string `json:"name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Password    string `json:"password,omitempty"`
}

func (u *User) Validate() *errors.RestErr {

	u.Name = strings.Trim(u.Name, " ")
	if len(u.Name) == 0 {
		return errors.NewBadRequestError("name can't be empty")
	}
	u.Password = strings.Trim(u.Password, " ")
	if len(u.Password) == 0 {
		return errors.NewBadRequestError("PASSWORD can't be empty")
	}
	u.Email = strings.Trim(u.Email, " ")
	if len(u.Email) == 0 {
		return errors.NewBadRequestError("email can't be empty")
	}

	return nil
}
