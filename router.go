package main

import (
	_ "main/docs"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type jwtCustomClaims struct {
	Account string `json:"account"`
	Email   string `json:"email"`
	jwt.StandardClaims
}

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/mainActivities", mainActivities) // POST /mainActivities
	e.POST("/signup", signup)                // POST /signup
	e.POST("/login", login)                  // POST /login
	e.GET("/", accessible)

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}))

	api.GET("", getAccountFromJWT)
	api.GET("/getAccountFromJWT", getAccountFromJWT) // GET /api/getAccountFromJWT

	return e
}
