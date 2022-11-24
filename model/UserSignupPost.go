package model

import (
	"errors"
	"fmt"
	"main/util"
)

type UserSignupPost struct {
	Account  string `db:"fAccount"`
	Password string `db:"fPassword"`
	Email    string `db:"fEmail"`
}

func CreateUser(sqlConnectionString string, u *UserSignupPost) (model UserSignupPost, err error) {
	var user UserSignupPost
	if u.Account == "" {
		outputErr := errors.New("Account empty")
		return user, outputErr
	}

	queryString := `INSERT INTO tMember (fAccount, fPassword, fEmail) VALUES (?,?,?)`
	rowId, result, err := util.SQLExec(sqlConnectionString, false, queryString, u.Account, u.Password, u.Email)
	if err != nil {
		return user, err
	}
	fmt.Println(rowId + result)
	user.Account = u.Account
	user.Password = u.Password
	user.Email = u.Email

	return user, nil
}
