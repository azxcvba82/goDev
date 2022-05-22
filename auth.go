package main

import (
	"main/model"
	"main/util"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var signingKey = []byte("secret")

// @Tags         Token
// @Description login user
// @Accept  json
// @Param login body object true "json"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /login [post]
func login(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello, World!")
	u := new(model.UserLoginPost)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.Account == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "empty name",
		}
	}
	user, err := model.FindUser(util.GetSQLConnectString(), &model.UserLoginPost{Account: u.Account})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	} else if user.Account != u.Account || user.Password != u.Password {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	claims := &jwtCustomClaims{
		user.Account,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// @Tags         Token
// @Description create user
// @Accept  json
// @Param signup body object true "json"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /signup [post]
func signup(c echo.Context) error {
	user := new(model.UserLoginPost)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Account == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	u, err := model.FindUser(util.GetSQLConnectString(), &model.UserLoginPost{Account: user.Account})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: err.Error(),
		}
	} else if u.Account == user.Account {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "name already exists",
		}
	}

	model.CreateUser(util.GetSQLConnectString(), user)
	return c.JSON(http.StatusCreated, user)
}

// @Tags         Token
// @Description getAccountFromJWT
// @Accept  json
// @Param user path string true "token"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/getAccountFromJWT [get]
// @security securityDefinitions.apikey ApiKeyAuth
func getAccountFromJWT(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account
	return c.JSON(http.StatusOK, account)
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
