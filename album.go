package main

import (
	"main/model"
	"main/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Tags         Album
// @Description getAlbumById load
// @Param id query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /getAlbumById [get]
func getAlbumById(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" || len(id) > 20 {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "parameter empty or not in scope",
		}
	}
	_, err := strconv.Atoi(id)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "id only allow number",
		}
	}
	album, err := model.GetAlbumById(util.GetSQLConnectString(), c.QueryParam("id"))
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, album)
}
