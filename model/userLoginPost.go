package model

import (
	"errors"
	"main/util"
)

type UserSSOLoginPost struct {
	StateBase64   string `db:"StateBase64"`
	IdTokenBase64 string `db:"IdTokenBase64"`
}

type UserVerifyPost struct {
	Token string `db:"token"`
}

func FindUser(sqlConnectionString string, u *UserSignupPost) (model UserSignupPost, err error) {
	var user UserSignupPost
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
