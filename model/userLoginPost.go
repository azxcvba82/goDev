package model

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type UserLoginPost struct {
	Account  string `string:"account"`
	Password string `string:"password"`
	Email    string `string:"email"`
}

func FindUser(u *UserLoginPost) UserLoginPost {
	var user UserLoginPost
	if u.Account == "" {
		return user
	}

	//os.Setenv("SQLCONNECTSTRING", "root:@tcp(20.99.156.107:3306)/godev")
	db, err := sql.Open("mysql", os.Getenv("SQLCONNECTSTRING"))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	nullfEmail := new(sql.NullString)
	queryString := `SELECT fAccount, fPassword, fEmail  FROM tMember WHERE fAccount = ? LIMIT 1`
	rows, err := db.Query(queryString, u.Account)
	checkErr(err)

	for rows.Next() {
		var fAccount string
		var fPassword string
		//var fEmail string
		err = rows.Scan(&fAccount, &fPassword, nullfEmail)

		checkErr(err)
		fmt.Println(fAccount)
		if fAccount == u.Account {
			user.Account = fAccount
			user.Password = fPassword
			if nullfEmail.Valid {
				user.Email = nullfEmail.String
			}
		}
	}
	return user
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
