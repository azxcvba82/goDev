package main

import (
	"main/model"
	"main/util"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
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

	album, err := model.GetAlbumById(util.GetSQLConnectStringRead(), id)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, album)
}

// @Tags         Album
// @Description getAlbumsByKindId load
// @Param kindId query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /getAlbumsByKindId [get]
func getAlbumsByKindId(c echo.Context) error {
	id := c.QueryParam("kindId")
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
	album, err := model.GetAlbumsByKindId(util.GetSQLConnectStringRead(), id)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, album)
}

// @Tags         Album
// @Description getProductsByAlbumId load
// @Param albumId query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /getProductsByAlbumId [get]
func getProductsByAlbumId(c echo.Context) error {
	id := c.QueryParam("albumId")
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
	album, err := model.GetProductsByAlbumId(util.GetSQLConnectStringRead(), id)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, album)
}

// @Tags         Album
// @Description getPlayListByAccount load
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/getPlayListByAccount [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func getPlayListByAccount(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account

	product, err := model.GetPlayListByAccount(util.GetSQLConnectStringRead(), account)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, product)
}

// @Tags         Album
// @Description addPlayLists load
// @Param productId query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/addPlayLists [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func addPlayLists(c echo.Context) error {
	productId := c.QueryParam("productId")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account

	product, err := model.UserAddPlayLists(util.GetSQLConnectString(), account, productId)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, product)
}

// @Tags         Album
// @Description deletePlayLists load
// @Param productId query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/deletePlayLists [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func deletePlayLists(c echo.Context) error {
	productId := c.QueryParam("productId")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account

	product, err := model.UserDeletePlayLists(util.GetSQLConnectString(), account, productId)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, product)
}

// @Tags         Album
// @Description addPlayLists load
// @Param productId query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/addPlayLists [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func addPlayLists(c echo.Context) error {
	productId := c.QueryParam("productId")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account

	product, err := model.UserAddPlayLists(util.GetSQLConnectString(), account, productId)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, product)
}

// @Tags         Album
// @Description deletePlayLists load
// @Param productId query string true "string valid"
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/deletePlayLists [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth
func deletePlayLists(c echo.Context) error {
	productId := c.QueryParam("productId")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account

	product, err := model.UserDeletePlayLists(util.GetSQLConnectString(), account, productId)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, product)
}
