package util

import (
	"fmt"
	"os"
)

func GetSQLConnectString() string {
	return os.Getenv("SQLCONNECTSTRING")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
