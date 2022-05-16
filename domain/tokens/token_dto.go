package tokens

import (
	"bookstore/token-jwt/utils/errors"
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Token struct {
	Id int64 `json:"id,omitempty"`
	UserId int64 `json:"userId,omitempty"`
	Token string `json:"token,omitempty"`
	Expires int64 `json:"expires,omitempty"`
}

func (t *Token) IsExpired() bool {
	return time.Unix(t.Expires, 0).Before(time.Now().UTC())
}

func (t *Token) Generate(userId int64) *errors.RestErr {
    b := make([]byte, 40)
    if _, err := rand.Read(b); err != nil {
        return errors.NewBadRequestError(err.Error())
    }
	t.Token = hex.EncodeToString(b)
	t.UserId = userId
	t.Expires = time.Now().UTC().Add(24 * time.Hour).Unix()

	err := t.Create()
	if err != nil {
		return err
	}

	return nil
}