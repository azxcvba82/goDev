package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
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
	return c.String(http.StatusOK, "Hello, World!")
}
