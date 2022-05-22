package model

import (
	"database/sql"
	"errors"
	"fmt"
	"main/util"
)

type UserLoginPost struct {
	Account  string `string:"account"`
	Password string `string:"password"`
	Email    string `string:"email"`
}

func FindUser(sqlConnectionString string, u *UserLoginPost) (model UserLoginPost, err error) {
	var user UserLoginPost
	if u.Account == "" {
		outputErr := errors.New("Account empty")
		return user, outputErr
	}

	nullfEmail := new(sql.NullString)
	queryString := `SELECT fAccount, fPassword, fEmail  FROM tMember WHERE fAccount = ? LIMIT 1`
	rows, err := util.SQLQuery(sqlConnectionString, queryString, u.Account)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		var fAccount string
		var fPassword string
		//var fEmail string
		err = rows.Scan(&fAccount, &fPassword, nullfEmail)

		util.CheckErr(err)
		fmt.Println(fAccount)
		if fAccount == u.Account {
			user.Account = fAccount
			user.Password = fPassword
			if nullfEmail.Valid {
				user.Email = nullfEmail.String
			}
		}
	}
	return user, nil
}

func CreateUser(sqlConnectionString string, u *UserLoginPost) (model UserLoginPost, err error) {
	var user UserLoginPost
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
