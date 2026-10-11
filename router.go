package main

import (
	"log"
	"main/config"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	_ = cfg

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/mainActivities", mainActivities)                     // GET /mainActivities
	e.GET("/mainAlbums", mainAlbums)                             // GET /mainAlbums
	e.GET("/getAlbumById", getAlbumById)                         // GET /getAlbumById
	e.GET("/getAlbumsByKindId", getAlbumsByKindId)               // GET /getAlbumsByKindId
	e.GET("/getProductsByAlbumId", getProductsByAlbumId)         // GET /getProductsByAlbumId
	e.GET("/getProductsByProductName", getProductsByProductName) // GET /getProductsByProductName

	e.GET("/allkind", allkind)           // GET /allkind
	e.GET("/allAlbumType", allAlbumType) // GET /allAlbumType

	authHandler := NewAuthHandler(cfg.JWTSecret)
	e.POST("/signup", authHandler.Signup)     // POST /signup
	e.POST("/verify", authHandler.Verify)     // POST /verify
	e.POST("/login", authHandler.Login)       // POST /login
	e.GET("/getSSOConfig", getSSOConfig)      // GET /getSSOConfig
	e.POST("/ssoLogin", authHandler.SsoLogin) // POST /ssoLogin
	e.GET("/", accessible)

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: cfg.JWTSecret,
	}))

	api.GET("", getAccountFromJWT)
	api.GET("/flushAllCache", flushAllCache)                       // GET /api/flushAllCache
	api.GET("/getAccountFromJWT", getAccountFromJWT)               // GET /api/getAccountFromJWT
	api.GET("/getShoppingCartByAccount", getShoppingCartByAccount) // GET /api/getShoppingCartByAccount
	api.GET("/getPlayListByAccount", getPlayListByAccount)         // GET /api/getPlayListByAccount
	api.GET("/addPlayLists", addPlayLists)                         // GET /api/addPlayLists
	api.GET("/deletePlayLists", deletePlayLists)                   // GET /api/deletePlayLists

	return e
}
