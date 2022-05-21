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

	e.POST("/signup", signup) // POST /signup
	e.POST("/login", login)   // POST /login
	e.GET("/", accessible)

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("secret"),
	}))

	api.GET("", test)
	api.GET("/test", test)                 // GET /api/todos
	api.POST("/todos", login)              // POST /api/todos
	api.DELETE("/todos/:id", login)        // DELETE /api/todos/:id
	api.PUT("/todos/:id/completed", login) // PUT /api/todos/:id/completed

	return e
}
