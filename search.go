package main

import (
	"main/model"
	"main/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags         Search
// @Description getProductsByProductName load
// @Param name query string false "string valid"
// @Param albumName query string false "string valid"
// @Param singer query string false "string valid"
// @Param group query string false "string valid"
// @Param composer query string false "string valid"
// @Param type query int false "int valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /getProductsByProductName [get]
func getProductsByProductName(c echo.Context) error {
	name := c.QueryParam("name")
	albumName := c.QueryParam("albumName")
	singer := c.QueryParam("singer")
	group := c.QueryParam("group")
	composer := c.QueryParam("composer")
	albumType := c.QueryParam("type")
	if len(name) > 40 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "parameter empty or not in scope",
		}
	}
	if name == "" && albumName == "" && singer == "" && group == "" && composer == "" && albumType == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "parameter empty or not in scope",
		}
	}
	item, err := model.GetProductsByProductName(util.GetSQLConnectString(), c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, item)
}

// @Tags         Search
// @Description allAlbumType load
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /allAlbumType [get]
func allAlbumType(c echo.Context) error {
	item, err := model.AllAlbumType(util.GetSQLConnectString())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, item)
}
