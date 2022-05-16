package users

// data access object

import (
	"bookstore/token-jwt/datasources/mysql/users_db"
	"bookstore/token-jwt/utils"
	"bookstore/token-jwt/utils/errors"
	"fmt"
	"strings"
	"time"
)

const (
	queryInsertUser             = "INSERT INTO users(name, email, passwordHash, createdAt) VALUES(?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, name, email status FROM users WHERE email=? AND passwordHash=?;"



)



func (u *User) Save() *errors.RestErr {

	if err := u.Validate(); err != nil {
		return err
	}
	u.DateCreated = time.Now().Format("2006-01-02T15:04:05Z")
	u.Password = utils.GetMd5(u.Password)
	


	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(u.Name, u.Email, u.Password, u.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	u.Id, err = insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	u.Password = ""
	return nil
}

func (u *User) FindByEmailAndPassword() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(u.Email, utils.GetMd5(u.Password))
	if err := result.Scan(&u.Id, &u.Name, &u.Email); err != nil {
		
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewBadRequestError("your email or password is wrong, Please try again with right ones")
		}

		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get u %d: %s", u.Id, err.Error()))
	}
	return nil
}
