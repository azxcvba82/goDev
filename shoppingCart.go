package main

import (
	"main/model"
	"main/util"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// @Tags         ShoppingCart
// @Description getShoppingCartByAccount load
// @Success 200 "ok"
// @Failure 500 "error"
// @Router /api/getShoppingCartByAccount [get]
// @security securityDefinitions.apikey BearerAuth
// @security BearerAuth

func getShoppingCartByAccount(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	account := claims.Account

	shoppingCart, err := model.GetShoppingCartByAccount(util.GetSQLConnectString(), account)

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, shoppingCart)
}
