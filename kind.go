package main

import (
	"main/model"
	"main/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Tags         Kind
// @Description allkind load
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /allkind [get]
func allkind(c echo.Context) error {
	kind, err := model.AllKind(util.GetSQLConnectStringRead())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, kind)
}
