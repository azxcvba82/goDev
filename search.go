package main

import (
	"main/model"
	"main/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags         Search
// @Description getProductsByProductName load
// @Param name query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /getProductsByProductName [get]
func getProductsByProductName(c echo.Context) error {
	id := c.QueryParam("name")
	if id == "" || len(id) > 40 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "parameter empty or not in scope",
		}
	}
	item, err := model.GetProductsByProductName(util.GetSQLConnectString(), id)
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
