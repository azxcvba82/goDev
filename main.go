package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"database/sql"
	_ "main/docs"

<<<<<<< HEAD
	_ "github.com/go-sql-driver/mysql"
=======
>>>>>>> 93caacb39fe58a4dbb51c8ea77697ee6a280241b
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)
	e.GET("/hello", hello)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}

// @Description get users in a group
// @Accept  json
// @Param group_id path int true "Group ID"
// @Param gender query string false "Gender" Enum(man, woman)
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /hello [get]
func hello(c echo.Context) error {
	//os.Setenv("SQLCONNECTSTRING", "root:@tcp(20.99.156.107:3306)/godev")
	db, err := sql.Open("mysql", os.Getenv("SQLCONNECTSTRING"))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)

	rows, err := db.Query("SELECT fAccount, fPassword, fEmail, fPrivilege  FROM tMember LIMIT 2")
	checkErr(err)

	mes := ""
	for rows.Next() {
		var fAccount string
		var fPassword string
		var fEmail string
		var fPrivilege int
		err = rows.Scan(&fAccount, &fPassword, &fEmail, &fPrivilege)
		checkErr(err)
		fmt.Println(fAccount)
		fmt.Println(fPassword)
		fmt.Println(fEmail)
		fmt.Println(fPrivilege)
		mes = fAccount + fPassword + fEmail
	}

	return c.String(http.StatusOK, "Hello, World!"+mes)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
