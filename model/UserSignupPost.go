package model

import (
	"errors"
	"fmt"
	"main/util"

	"gopkg.in/guregu/null.v4"
)

type UserSignupPost struct {
	Account  string      `db:"fAccount"`
	Password string      `db:"fPassword"`
	Email    null.String `db:"fEmail"`
}

func CreateUser(sqlConnectionString string, u *UserSignupPost) (model UserSignupPost, err error) {
	var user UserSignupPost
	if u.Account == "" {
		outputErr := errors.New("Account empty")
		return user, outputErr
	}

	queryString := `INSERT INTO tMember (fAccount, fPassword, fEmail) VALUES (?,?,?)`
	result, err := util.SQLExec(sqlConnectionString, false, queryString, u.Account, u.Password, u.Email)
	if err != nil {
		return user, err
	}
	fmt.Println(result)

	return user, nil
}
