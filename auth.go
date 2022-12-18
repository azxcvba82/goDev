package main

import (
	"context"
	"errors"
	"fmt"
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
	if err != nil && err.Error() != "sql: no rows in result set" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	} else if u.Account == user.Account {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
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
		return c.JSON(http.StatusBadRequest, err)
	}
	message := "https://azxcvba99.net/verify?token=" + t
	res, err := model.SendMail(util.GetSQLConnectStringRead(), user.Email, "Email verification", message)
	if err != nil {
		return c.JSON(http.StatusBadRequest, res+"/"+err.Error())
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

// @Tags         Token
// @Description verify account regist token
// @Accept  json
// @Param verify body object true "json"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /verify [post]
func verify(c echo.Context) error {

	u := new(model.UserVerifyPost)

	if err := c.Bind(u); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	if u.Token == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "empty token",
		}
	}

	tokenVerify, err := jwt.Parse(u.Token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if claims, ok := tokenVerify.Claims.(jwt.MapClaims); ok && tokenVerify.Valid {
		fmt.Printf("user_id: %v\n", string(claims["Account"].(string)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
		return c.JSON(http.StatusOK, map[string]string{
			"account": string(claims["Account"].(string)),
		})
	} else {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
