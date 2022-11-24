package main

import (
	"context"
	"errors"
	"main/model"
	"main/util"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
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
	u := new(model.UserSignupPost)
	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if u.Account == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "empty name",
		}
	}
	var userForCheck model.UserSignupPost
	user, err := model.FindUser(util.GetSQLConnectStringRead(), &model.UserSignupPost{Account: u.Account})
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	} else if user == userForCheck {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		}
	} else if user.Account != u.Account || user.Password != u.Password {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}
	expiresTime := time.Now().UTC().Add(time.Hour * 6)
	claims := &jwtCustomClaims{
		user.Account,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"account":   user.Account,
		"token":     t,
		"expiresAt": expiresTime.Format(time.RFC1123),
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
	user := new(model.UserSignupPost)
	if err := c.Bind(user); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	if user.Account == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	u, err := model.FindUser(util.GetSQLConnectStringRead(), &model.UserSignupPost{Account: user.Account})
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

	// _, err = model.CreateUser(util.GetSQLConnectString(), user)
	// if err != nil {
	// 	return &echo.HTTPError{
	// 		Code:    http.StatusBadRequest,
	// 		Message: err.Error(),
	// 	}
	// }
	expiresTime := time.Now().UTC().Add(time.Minute * 15)
	claims := &jwtCustomClaims{
		user.Account,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}
	message := "https://azxcvba99.net/verify?token=" + t
	res, err := model.SendMail(util.GetSQLConnectStringRead(), user.Email, "Email verification", message)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, res)
}

// @Tags         Token
// @Description getAccountFromJWT
// @Accept  json
// @Param user path string true "token"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/getAccountFromJWT [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func getAccountFromJWT(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account
	return c.JSON(http.StatusOK, account)
}

// @Tags         Token
// @Description getSSOConfig
// @Accept  json
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /getSSOConfig [get]
func getSSOConfig(c echo.Context) error {
	params := url.Values{}
	params.Add("state", "google")
	params.Add("response_type", "id_token")
	params.Add("nonce", time.Now().UTC().String())
	params.Add("response_mode", "fragment")
	params.Add("prompt", "select_account")
	params.Add("scope", "openid email profile")
	params.Add("client_id", "248375232247-fppbkkm4d0vjr8na12j64402nh9muu74.apps.googleusercontent.com")
	params.Add("redirect_uri", "https://azxcvba99.net/")
	//params.Add("redirect_uri", "http://localhost:3000/")

	var fullUrl string
	var baseUrl string
	baseUrl = "https://accounts.google.com/o/oauth2/auth?"
	fullUrl = baseUrl + params.Encode()
	return c.JSON(http.StatusOK, fullUrl)
}

// @Tags         Token
// @Description ssoLogin
// @Accept  json
// @Param ssoLogin body object true "json"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /ssoLogin [post]
func ssoLogin(c echo.Context) error {
	u := new(model.UserSSOLoginPost)
	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	payload, err := idtoken.Validate(context.Background(), u.IdTokenBase64, "248375232247-fppbkkm4d0vjr8na12j64402nh9muu74.apps.googleusercontent.com")
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error() + "sso test",
		}
	}

	if payload.Claims["email_verified"].(bool) == false {
		return errors.New("email not verified")
	}

	var userForCheck model.UserSignupPost
	user, err := model.FindUser(util.GetSQLConnectStringRead(), &model.UserSignupPost{Account: payload.Claims["email"].(string)})
	if user == userForCheck && err.Error() == "sql: no rows in result set" {
		userCreate, err := model.CreateUser(util.GetSQLConnectString(), &model.UserSignupPost{Account: payload.Claims["email"].(string), Password: "sso", Email: payload.Claims["email"].(string)})
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}
		}

		expiresTime := time.Now().UTC().Add(time.Hour * 6)
		claims := &jwtCustomClaims{
			userCreate.Account,
			userCreate.Email,
			jwt.StandardClaims{
				ExpiresAt: expiresTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString(signingKey)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"account":   userCreate.Account,
			"token":     t,
			"expiresAt": expiresTime.Format(time.RFC1123),
		})
	} else if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	expiresTime := time.Now().UTC().Add(time.Hour * 6)
	claims := &jwtCustomClaims{
		user.Account,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"account":   user.Account,
		"token":     t,
		"expiresAt": expiresTime.Format(time.RFC1123),
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
