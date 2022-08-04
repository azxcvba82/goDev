package model

import (
	"errors"
	"main/util"

	"gopkg.in/guregu/null.v4"
)

type UserLoginPost struct {
	Account  string      `db:"fAccount"`
	Password string      `db:"fPassword"`
	Email    null.String `db:"fEmail"`
}

func FindUser(sqlConnectionString string, u *UserLoginPost) (model UserLoginPost, err error) {
	var user UserLoginPost
	if u.Account == "" {
		outputErr := errors.New("Account empty")
		return user, outputErr
	}

	queryString := `SELECT fAccount, fPassword, fEmail  FROM tMember WHERE fAccount = ? LIMIT 1`
	err = util.SQLQueryV2(&user, sqlConnectionString, false, queryString, u.Account)
	if err != nil {
		return user, err
	}

	return user, nil
}
