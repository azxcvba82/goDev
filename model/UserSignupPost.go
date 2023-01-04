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

func CreateUser(sqlConnectionString string, u *UserSignupPost, isActive bool) (model UserSignupPost, err error) {
	var user UserSignupPost
	if u.Account == "" {
		outputErr := errors.New("Account empty")
		return user, outputErr
	}

	active := "N"
	if isActive == true {
		active = "Y"
	}

	queryString := `INSERT INTO tMember (fAccount, fPassword, fEmail, fisActive) VALUES (?,?,?,?)`
	rowId, result, err := util.SQLExec(sqlConnectionString, false, queryString, u.Account, u.Password, u.Email, active)
	if err != nil {
		return user, err
	}
	fmt.Println(rowId + result)
	user.Account = u.Account
	user.Password = u.Password
	user.Email = u.Email

	return user, nil
}

func ActivateUser(sqlConnectionString string, Account string) (err error) {

	queryString := `UPDATE tMember SET fisActive = 'Y' WHERE fAccount = ? `
	_, result, err := util.SQLExec(sqlConnectionString, false, queryString, Account)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}
