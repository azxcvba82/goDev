package model

import (
	"errors"
	"fmt"
	"main/util"
)

type UserSignupPost struct {
	Account  string `string:"account"`
	Password string `string:"password"`
	Email    string `string:"email"`
}

func CreateUser(sqlConnectionString string, u *UserSignupPost) (model UserSignupPost, err error) {
	var user UserSignupPost
	if u.Account == "" {
		outputErr := errors.New("Account empty")
		return user, outputErr
	}

	queryString := `INSERT INTO tMember (fAccount, fPassword, fEmail) VALUES (?,?,?)`
	result, err := util.SQLExec(sqlConnectionString, queryString, u.Account, u.Password, u.Email)
	if err != nil {
		return user, err
	}
	fmt.Println(result)

	return user, nil
}
