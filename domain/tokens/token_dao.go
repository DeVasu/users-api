package tokens

import (
	"bookstore/token-jwt/datasources/mysql/users_db"
	"bookstore/token-jwt/utils/errors"
	"fmt"
	"strings"
)

const (
	queryInsertToken string = "INSERT INTO tokens(userId, token, expires) VALUES(?, ?, ?);"
	queryGetTokenByUserId string = "SELECT id, token, expires from tokens WHERE userId=?;"
	queryGetUserByToken string = "SELECT userId, expires from tokens WHERE token=?;"
	queryDeleteToken string = "DELETE FROM tokens WHERE token=?;"
)

func (t* Token) Create() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertToken)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	insertRes, err := stmt.Exec(t.UserId, t.Token, t.Expires)
	if err != nil {
		return errors.NewInternalServerError("error when trhying to sabe token " + err.Error())
	}
	t.Id, err = insertRes.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when retrieving id of the new token")
	}

	return nil
}

func (t* Token) GetByUserId() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetTokenByUserId)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()	
	res := stmt.QueryRow(t.UserId)
	if err := res.Scan(&t.Id, &t.Token, &t.Expires); err != nil {

		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewInternalServerError("token does not exists for this user")
		}
		return errors.NewInternalServerError("error readin the user" + err.Error())
	}

	return nil
}

func (t *Token) GetUserByToken() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUserByToken)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow(t.Token)

	if err := res.Scan(&t.UserId, &t.Expires); err != nil {

		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewBadRequestError("token is invalid as no users match this token")
		}

		return errors.NewBadRequestError("user with this token does not exists")
	}

	return nil
}
func (t *Token) DeleteToken() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryDeleteToken)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	

	if _, err = stmt.Exec(t.Token); err != nil {
		fmt.Println(err.Error())

		return errors.NewInternalServerError("error when tying to delete User user" + err.Error())
	}

	return nil
}