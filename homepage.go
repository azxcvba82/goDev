package main

import (
	"main/model"
	"main/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags         Homepage
// @Description mainActivities load
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /mainActivities [get]
func mainActivities(c echo.Context) error {
	event, err := model.EventQuery(util.GetSQLConnectStringRead(), "")
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, event)
}

// @Tags         Homepage
// @Description mainAlbums load
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /mainAlbums [get]
func mainAlbums(c echo.Context) error {
	album, err := model.AllAlbum(util.GetSQLConnectStringRead())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, album)
}

// @Tags         Homepage
// @Description Flush All Cache
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/flushAllCache [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func flushAllCache(c echo.Context) error {
	err := util.FlushAllCache()
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, "done")
}
